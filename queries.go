package pqstream

var (
	sqlQueryTables = `
SELECT table_name
  FROM information_schema.tables
 WHERE table_schema='public'
   AND table_type='BASE TABLE'
`
	sqlTriggerFunction = `
CREATE OR REPLACE FUNCTION pqstream_notify() RETURNS TRIGGER AS $$
    DECLARE 
        payload json;
        notification json;
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            payload = row_to_json(OLD);
        ELSE
            payload = row_to_json(NEW);
        END IF;
        
        notification = json_build_object(
                          'schema', TG_TABLE_SCHEMA,
                          'table', TG_TABLE_NAME,
                          'op', TG_OP,
						  'id', json_extract_path(payload, 'id')::text,
                          'payload', payload);
        IF (length(notification::text) > 8000) THEN
          notification = json_build_object(
                            'schema', TG_TABLE_SCHEMA,
                            'table', TG_TABLE_NAME,
                            'op', TG_OP,
							'id', json_extract_path(payload, 'id')::text);
        END IF;
        
        PERFORM pg_notify('pqstream_notify', notification::text);
        RETURN NULL; 
    END;
$$ LANGUAGE plpgsql;
`
	sqlRemoveTrigger = `
DROP TRIGGER IF EXISTS pqstream_notify ON %s
`
	sqlInstallTrigger = `
CREATE TRIGGER pqstream_notify
AFTER INSERT OR UPDATE OR DELETE ON %s
    FOR EACH ROW EXECUTE PROCEDURE pqstream_notify();
`
	sqlFetchRowById = `
	SELECT row_to_json(r)::text from (select * from %s where id = $1::%s) r;
`
)
