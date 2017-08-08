// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timeseries.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	timeseries.proto

It has these top-level messages:
	TSResponse
	Point
	Query
	Meta
	Tag
	MetaFound
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type TSResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *TSResponse) Reset()                    { *m = TSResponse{} }
func (m *TSResponse) String() string            { return proto1.CompactTextString(m) }
func (*TSResponse) ProtoMessage()               {}
func (*TSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TSResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type Point struct {
	Ksid  string  `protobuf:"bytes,1,opt,name=ksid" json:"ksid,omitempty"`
	Tsid  string  `protobuf:"bytes,2,opt,name=tsid" json:"tsid,omitempty"`
	Value float32 `protobuf:"fixed32,3,opt,name=value" json:"value,omitempty"`
	Date  int64   `protobuf:"varint,4,opt,name=date" json:"date,omitempty"`
	Empty bool    `protobuf:"varint,5,opt,name=empty" json:"empty,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto1.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Point) GetKsid() string {
	if m != nil {
		return m.Ksid
	}
	return ""
}

func (m *Point) GetTsid() string {
	if m != nil {
		return m.Tsid
	}
	return ""
}

func (m *Point) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Point) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Point) GetEmpty() bool {
	if m != nil {
		return m.Empty
	}
	return false
}

type Query struct {
	Ksid  string `protobuf:"bytes,1,opt,name=ksid" json:"ksid,omitempty"`
	Tsid  string `protobuf:"bytes,2,opt,name=tsid" json:"tsid,omitempty"`
	Start int64  `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	End   int64  `protobuf:"varint,4,opt,name=end" json:"end,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto1.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Query) GetKsid() string {
	if m != nil {
		return m.Ksid
	}
	return ""
}

func (m *Query) GetTsid() string {
	if m != nil {
		return m.Tsid
	}
	return ""
}

func (m *Query) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Query) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type Meta struct {
	Ksid   string `protobuf:"bytes,1,opt,name=ksid" json:"ksid,omitempty"`
	Tsid   string `protobuf:"bytes,2,opt,name=tsid" json:"tsid,omitempty"`
	Metric string `protobuf:"bytes,3,opt,name=metric" json:"metric,omitempty"`
	Tags   []*Tag `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Meta) Reset()                    { *m = Meta{} }
func (m *Meta) String() string            { return proto1.CompactTextString(m) }
func (*Meta) ProtoMessage()               {}
func (*Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Meta) GetKsid() string {
	if m != nil {
		return m.Ksid
	}
	return ""
}

func (m *Meta) GetTsid() string {
	if m != nil {
		return m.Tsid
	}
	return ""
}

func (m *Meta) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *Meta) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto1.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MetaFound struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *MetaFound) Reset()                    { *m = MetaFound{} }
func (m *MetaFound) String() string            { return proto1.CompactTextString(m) }
func (*MetaFound) ProtoMessage()               {}
func (*MetaFound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MetaFound) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto1.RegisterType((*TSResponse)(nil), "proto.TSResponse")
	proto1.RegisterType((*Point)(nil), "proto.Point")
	proto1.RegisterType((*Query)(nil), "proto.Query")
	proto1.RegisterType((*Meta)(nil), "proto.Meta")
	proto1.RegisterType((*Tag)(nil), "proto.Tag")
	proto1.RegisterType((*MetaFound)(nil), "proto.MetaFound")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Timeseries service

type TimeseriesClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (Timeseries_WriteClient, error)
	Read(ctx context.Context, in *Query, opts ...grpc.CallOption) (Timeseries_ReadClient, error)
	GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*MetaFound, error)
}

type timeseriesClient struct {
	cc *grpc.ClientConn
}

func NewTimeseriesClient(cc *grpc.ClientConn) TimeseriesClient {
	return &timeseriesClient{cc}
}

func (c *timeseriesClient) Write(ctx context.Context, opts ...grpc.CallOption) (Timeseries_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Timeseries_serviceDesc.Streams[0], c.cc, "/proto.Timeseries/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeseriesWriteClient{stream}
	return x, nil
}

type Timeseries_WriteClient interface {
	Send(*Point) error
	CloseAndRecv() (*TSResponse, error)
	grpc.ClientStream
}

type timeseriesWriteClient struct {
	grpc.ClientStream
}

func (x *timeseriesWriteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *timeseriesWriteClient) CloseAndRecv() (*TSResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TSResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *timeseriesClient) Read(ctx context.Context, in *Query, opts ...grpc.CallOption) (Timeseries_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Timeseries_serviceDesc.Streams[1], c.cc, "/proto.Timeseries/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeseriesReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Timeseries_ReadClient interface {
	Recv() (*Point, error)
	grpc.ClientStream
}

type timeseriesReadClient struct {
	grpc.ClientStream
}

func (x *timeseriesReadClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *timeseriesClient) GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*MetaFound, error) {
	out := new(MetaFound)
	err := grpc.Invoke(ctx, "/proto.Timeseries/GetMeta", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Timeseries service

type TimeseriesServer interface {
	Write(Timeseries_WriteServer) error
	Read(*Query, Timeseries_ReadServer) error
	GetMeta(context.Context, *Meta) (*MetaFound, error)
}

func RegisterTimeseriesServer(s *grpc.Server, srv TimeseriesServer) {
	s.RegisterService(&_Timeseries_serviceDesc, srv)
}

func _Timeseries_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TimeseriesServer).Write(&timeseriesWriteServer{stream})
}

type Timeseries_WriteServer interface {
	SendAndClose(*TSResponse) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type timeseriesWriteServer struct {
	grpc.ServerStream
}

func (x *timeseriesWriteServer) SendAndClose(m *TSResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *timeseriesWriteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Timeseries_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeseriesServer).Read(m, &timeseriesReadServer{stream})
}

type Timeseries_ReadServer interface {
	Send(*Point) error
	grpc.ServerStream
}

type timeseriesReadServer struct {
	grpc.ServerStream
}

func (x *timeseriesReadServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

func _Timeseries_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeseriesServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Timeseries/GetMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeseriesServer).GetMeta(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timeseries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Timeseries",
	HandlerType: (*TimeseriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeta",
			Handler:    _Timeseries_GetMeta_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Timeseries_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _Timeseries_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timeseries.proto",
}

func init() { proto1.RegisterFile("timeseries.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0xed, 0xe6, 0xa3, 0xf7, 0x76, 0x7a, 0xb9, 0xd4, 0x45, 0x24, 0x54, 0x91, 0xb0, 0x0f, 0x12,
	0x44, 0x8b, 0xd4, 0xff, 0xa0, 0x4f, 0x82, 0xae, 0x81, 0x3e, 0xaf, 0x66, 0x2c, 0x21, 0x36, 0x5b,
	0xb2, 0x53, 0xa1, 0x7f, 0xc2, 0xdf, 0x2c, 0x3b, 0x1b, 0xeb, 0xc7, 0x93, 0x3e, 0xe5, 0x9c, 0xb3,
	0x93, 0x39, 0x87, 0x33, 0x30, 0xa1, 0x7a, 0x85, 0x0e, 0xbb, 0x1a, 0xdd, 0x6c, 0xdd, 0x59, 0xb2,
	0x32, 0xe5, 0x8f, 0x3a, 0x02, 0x28, 0xef, 0x35, 0xba, 0xb5, 0x6d, 0x1d, 0xca, 0xff, 0x10, 0xd9,
	0x26, 0x13, 0xb9, 0x28, 0xfe, 0xea, 0xc8, 0x36, 0xca, 0x42, 0x7a, 0x6b, 0xeb, 0x96, 0xa4, 0x84,
	0xa4, 0x71, 0x75, 0xc5, 0x4f, 0x23, 0xcd, 0xd8, 0x6b, 0xe4, 0xb5, 0x28, 0x68, 0x1e, 0xcb, 0x7d,
	0x48, 0x5f, 0xcc, 0xf3, 0x06, 0xb3, 0x38, 0x17, 0x45, 0xa4, 0x03, 0xf1, 0x93, 0x95, 0x21, 0xcc,
	0x92, 0x5c, 0x14, 0xb1, 0x66, 0xec, 0x27, 0x71, 0xb5, 0xa6, 0x6d, 0x96, 0xb2, 0x5b, 0x20, 0x6a,
	0x01, 0xe9, 0xdd, 0x06, 0xbb, 0xed, 0x6f, 0x0c, 0x1d, 0x99, 0x8e, 0xd8, 0x30, 0xd6, 0x81, 0xc8,
	0x09, 0xc4, 0xd8, 0x56, 0xbd, 0x9f, 0x87, 0xea, 0x09, 0x92, 0x1b, 0x24, 0xf3, 0xe3, 0xbd, 0x07,
	0x30, 0x5c, 0x21, 0x75, 0xf5, 0x23, 0x2f, 0x1e, 0xe9, 0x9e, 0xc9, 0x63, 0x48, 0xc8, 0x2c, 0x5d,
	0x96, 0xe4, 0x71, 0x31, 0x9e, 0x43, 0x28, 0x73, 0x56, 0x9a, 0xa5, 0x66, 0x5d, 0x9d, 0x43, 0x5c,
	0x9a, 0xa5, 0x0f, 0xd0, 0xe0, 0xb6, 0x77, 0xf1, 0xf0, 0xa3, 0x99, 0xe0, 0x12, 0x88, 0x3a, 0x84,
	0x91, 0x8f, 0x75, 0x65, 0x37, 0x6d, 0xf5, 0xbd, 0xfd, 0xf9, 0xab, 0x00, 0x28, 0x77, 0x77, 0x93,
	0x67, 0x90, 0x2e, 0xba, 0x9a, 0x50, 0xfe, 0xeb, 0x5d, 0xf9, 0x34, 0xd3, 0xbd, 0xf7, 0x0c, 0xbb,
	0x33, 0xaa, 0x41, 0x21, 0xe4, 0x09, 0x24, 0x1a, 0x4d, 0xb5, 0x1b, 0xe6, 0x5a, 0xa7, 0x5f, 0x7e,
	0x55, 0x83, 0x0b, 0x21, 0x4f, 0xe1, 0xcf, 0x35, 0x12, 0x77, 0x33, 0xee, 0x1f, 0x3d, 0x99, 0x4e,
	0x3e, 0x11, 0x8e, 0xa7, 0x06, 0x0f, 0x43, 0x96, 0x2e, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57,
	0xfb, 0xe1, 0x01, 0x4e, 0x02, 0x00, 0x00,
}
