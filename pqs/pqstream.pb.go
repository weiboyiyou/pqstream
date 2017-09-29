// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pqstream.proto

/*
Package pqs is a generated protocol buffer package.

It is generated from these files:
	pqstream.proto

It has these top-level messages:
	ListenRequest
	RawEvent
	Event
*/
package pqs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An operation in the database.
type Operation int32

const (
	Operation_UNKNOWN  Operation = 0
	Operation_INSERT   Operation = 1
	Operation_UPDATE   Operation = 2
	Operation_DELETE   Operation = 3
	Operation_TRUNCATE Operation = 4
)

var Operation_name = map[int32]string{
	0: "UNKNOWN",
	1: "INSERT",
	2: "UPDATE",
	3: "DELETE",
	4: "TRUNCATE",
}
var Operation_value = map[string]int32{
	"UNKNOWN":  0,
	"INSERT":   1,
	"UPDATE":   2,
	"DELETE":   3,
	"TRUNCATE": 4,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A request to listen to database event streams.
type ListenRequest struct {
	// if provided, this string will be used to match table names to track.
	TableRegexp string `protobuf:"bytes,1,opt,name=table_regexp,json=tableRegexp" json:"table_regexp,omitempty"`
}

func (m *ListenRequest) Reset()                    { *m = ListenRequest{} }
func (m *ListenRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()               {}
func (*ListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenRequest) GetTableRegexp() string {
	if m != nil {
		return m.TableRegexp
	}
	return ""
}

// RawEvent is an internal type.
type RawEvent struct {
	Schema   string                  `protobuf:"bytes,1,opt,name=schema" json:"schema,omitempty"`
	Table    string                  `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	Op       Operation               `protobuf:"varint,3,opt,name=op,enum=pqs.Operation" json:"op,omitempty"`
	Id       string                  `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
	Payload  *google_protobuf.Struct `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
	Previous *google_protobuf.Struct `protobuf:"bytes,6,opt,name=previous" json:"previous,omitempty"`
}

func (m *RawEvent) Reset()                    { *m = RawEvent{} }
func (m *RawEvent) String() string            { return proto.CompactTextString(m) }
func (*RawEvent) ProtoMessage()               {}
func (*RawEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RawEvent) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *RawEvent) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *RawEvent) GetOp() Operation {
	if m != nil {
		return m.Op
	}
	return Operation_UNKNOWN
}

func (m *RawEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RawEvent) GetPayload() *google_protobuf.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RawEvent) GetPrevious() *google_protobuf.Struct {
	if m != nil {
		return m.Previous
	}
	return nil
}

// A database event.
type Event struct {
	Schema string    `protobuf:"bytes,1,opt,name=schema" json:"schema,omitempty"`
	Table  string    `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	Op     Operation `protobuf:"varint,3,opt,name=op,enum=pqs.Operation" json:"op,omitempty"`
	// if the id column exists, this will populate it
	Id string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
	// payload is a json encoded representation of the changed object.
	Payload *google_protobuf.Struct `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
	// patch is, in the event of op==UPDATE an RFC7386 JSON merge patch.
	Patch *google_protobuf.Struct `protobuf:"bytes,6,opt,name=patch" json:"patch,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Event) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Event) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Event) GetOp() Operation {
	if m != nil {
		return m.Op
	}
	return Operation_UNKNOWN
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetPayload() *google_protobuf.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetPatch() *google_protobuf.Struct {
	if m != nil {
		return m.Patch
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenRequest)(nil), "pqs.ListenRequest")
	proto.RegisterType((*RawEvent)(nil), "pqs.RawEvent")
	proto.RegisterType((*Event)(nil), "pqs.Event")
	proto.RegisterEnum("pqs.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PQStream service

type PQStreamClient interface {
	// Listen responds with a stream of database operations.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (PQStream_ListenClient, error)
}

type pQStreamClient struct {
	cc *grpc.ClientConn
}

func NewPQStreamClient(cc *grpc.ClientConn) PQStreamClient {
	return &pQStreamClient{cc}
}

func (c *pQStreamClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (PQStream_ListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PQStream_serviceDesc.Streams[0], c.cc, "/pqs.PQStream/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &pQStreamListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PQStream_ListenClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type pQStreamListenClient struct {
	grpc.ClientStream
}

func (x *pQStreamListenClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PQStream service

type PQStreamServer interface {
	// Listen responds with a stream of database operations.
	Listen(*ListenRequest, PQStream_ListenServer) error
}

func RegisterPQStreamServer(s *grpc.Server, srv PQStreamServer) {
	s.RegisterService(&_PQStream_serviceDesc, srv)
}

func _PQStream_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PQStreamServer).Listen(m, &pQStreamListenServer{stream})
}

type PQStream_ListenServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type pQStreamListenServer struct {
	grpc.ServerStream
}

func (x *pQStreamListenServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _PQStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pqs.PQStream",
	HandlerType: (*PQStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _PQStream_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pqstream.proto",
}

func init() { proto.RegisterFile("pqstream.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xc1, 0xaf, 0x93, 0x40,
	0x10, 0xc6, 0xdf, 0xd2, 0x07, 0x8f, 0x37, 0xad, 0x84, 0x6c, 0x8c, 0x92, 0x1e, 0x4c, 0xe5, 0xd4,
	0x18, 0x05, 0xa5, 0x31, 0xf1, 0x6a, 0x2c, 0x07, 0xb5, 0xa1, 0x75, 0x4b, 0xe3, 0xd1, 0x2c, 0x74,
	0xa5, 0x24, 0x94, 0x5d, 0xd8, 0xa5, 0xda, 0xbf, 0xd0, 0xab, 0x7f, 0x92, 0xe9, 0xd2, 0x36, 0xf1,
	0xa2, 0xd7, 0x77, 0xda, 0x99, 0x6f, 0xbe, 0x2f, 0x99, 0xdf, 0x66, 0xc0, 0x11, 0x8d, 0x54, 0x2d,
	0xa3, 0xfb, 0x40, 0xb4, 0x5c, 0x71, 0x3c, 0x10, 0x8d, 0x1c, 0xbf, 0x2d, 0x4a, 0xb5, 0xeb, 0xb2,
	0x20, 0xe7, 0xfb, 0xb0, 0xe0, 0x15, 0xad, 0x8b, 0x50, 0x4f, 0xb3, 0xee, 0x7b, 0x28, 0xd4, 0x51,
	0x30, 0x19, 0x4a, 0xd5, 0x76, 0xb9, 0x3a, 0x3f, 0x7d, 0xd6, 0x8f, 0xe0, 0xd1, 0xa2, 0x94, 0x8a,
	0xd5, 0x84, 0x35, 0x1d, 0x93, 0x0a, 0x3f, 0x87, 0x91, 0xa2, 0x59, 0xc5, 0xbe, 0xb5, 0xac, 0x60,
	0x3f, 0x85, 0x87, 0x26, 0x68, 0x7a, 0x4f, 0x86, 0x5a, 0x23, 0x5a, 0xf2, 0x7f, 0x23, 0xb0, 0x09,
	0xfd, 0x11, 0x1f, 0x58, 0xad, 0xf0, 0x13, 0xb0, 0x64, 0xbe, 0x63, 0x7b, 0x7a, 0x76, 0x9e, 0x3b,
	0xfc, 0x18, 0x4c, 0x9d, 0xf1, 0x0c, 0x2d, 0xf7, 0x0d, 0x7e, 0x06, 0x06, 0x17, 0xde, 0x60, 0x82,
	0xa6, 0x4e, 0xe4, 0x04, 0xa2, 0x91, 0xc1, 0x52, 0xb0, 0x96, 0xaa, 0x92, 0xd7, 0xc4, 0xe0, 0x02,
	0x3b, 0x60, 0x94, 0x5b, 0xef, 0x56, 0x47, 0x8c, 0x72, 0x8b, 0xdf, 0xc0, 0x9d, 0xa0, 0xc7, 0x8a,
	0xd3, 0xad, 0x67, 0x4e, 0xd0, 0x74, 0x18, 0x3d, 0x0d, 0x0a, 0xce, 0x8b, 0x8a, 0x05, 0x17, 0xb8,
	0x60, 0xad, 0x71, 0xc8, 0xc5, 0x87, 0x67, 0x60, 0x8b, 0x96, 0x1d, 0x4a, 0xde, 0x49, 0xcf, 0xfa,
	0x77, 0xe6, 0x6a, 0xf4, 0x7f, 0x21, 0x30, 0x1f, 0x28, 0xcf, 0x2b, 0x30, 0x05, 0x55, 0xf9, 0xee,
	0x7f, 0x30, 0xbd, 0xeb, 0xc5, 0x27, 0xb8, 0xbf, 0xae, 0x80, 0x87, 0x70, 0xb7, 0x49, 0x3e, 0x27,
	0xcb, 0xaf, 0x89, 0x7b, 0x83, 0x01, 0xac, 0x8f, 0xc9, 0x3a, 0x26, 0xa9, 0x8b, 0x4e, 0xf5, 0x66,
	0x35, 0x7f, 0x9f, 0xc6, 0xae, 0x71, 0xaa, 0xe7, 0xf1, 0x22, 0x4e, 0x63, 0x77, 0x80, 0x47, 0x60,
	0xa7, 0x64, 0x93, 0x7c, 0x38, 0x4d, 0x6e, 0xa3, 0x77, 0x60, 0xaf, 0xbe, 0xac, 0xf5, 0xa9, 0xe1,
	0x97, 0x60, 0xf5, 0x87, 0x82, 0xb1, 0xe6, 0xfc, 0xeb, 0x6a, 0xc6, 0xa0, 0x35, 0xfd, 0x83, 0xfe,
	0xcd, 0x6b, 0x94, 0x59, 0x7a, 0xbb, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xed, 0xd0,
	0x65, 0xab, 0x02, 0x00, 0x00,
}
