// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package com is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	ComMsg
	StatusReport
*/
package com

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type StatusReport_Status int32

const (
	StatusReport_SUCCESS StatusReport_Status = 0
	StatusReport_ERROR   StatusReport_Status = 1
)

var StatusReport_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
}
var StatusReport_Status_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
}

func (x StatusReport_Status) String() string {
	return proto.EnumName(StatusReport_Status_name, int32(x))
}
func (StatusReport_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ComMsg struct {
	Msg    string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Sender int32  `protobuf:"varint,2,opt,name=sender" json:"sender,omitempty"`
}

func (m *ComMsg) Reset()                    { *m = ComMsg{} }
func (m *ComMsg) String() string            { return proto.CompactTextString(m) }
func (*ComMsg) ProtoMessage()               {}
func (*ComMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ComMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ComMsg) GetSender() int32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

type StatusReport struct {
	Status  StatusReport_Status `protobuf:"varint,1,opt,name=status,enum=com.StatusReport_Status" json:"status,omitempty"`
	Message *ComMsg             `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StatusReport) Reset()                    { *m = StatusReport{} }
func (m *StatusReport) String() string            { return proto.CompactTextString(m) }
func (*StatusReport) ProtoMessage()               {}
func (*StatusReport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusReport) GetStatus() StatusReport_Status {
	if m != nil {
		return m.Status
	}
	return StatusReport_SUCCESS
}

func (m *StatusReport) GetMessage() *ComMsg {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*ComMsg)(nil), "com.ComMsg")
	proto.RegisterType((*StatusReport)(nil), "com.StatusReport")
	proto.RegisterEnum("com.StatusReport_Status", StatusReport_Status_name, StatusReport_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteMsg service

type RouteMsgClient interface {
	// Ask the consumer to consume a MSG, wait for a response
	SendMsg(ctx context.Context, in *ComMsg, opts ...grpc.CallOption) (*StatusReport, error)
}

type routeMsgClient struct {
	cc *grpc.ClientConn
}

func NewRouteMsgClient(cc *grpc.ClientConn) RouteMsgClient {
	return &routeMsgClient{cc}
}

func (c *routeMsgClient) SendMsg(ctx context.Context, in *ComMsg, opts ...grpc.CallOption) (*StatusReport, error) {
	out := new(StatusReport)
	err := grpc.Invoke(ctx, "/com.RouteMsg/SendMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RouteMsg service

type RouteMsgServer interface {
	// Ask the consumer to consume a MSG, wait for a response
	SendMsg(context.Context, *ComMsg) (*StatusReport, error)
}

func RegisterRouteMsgServer(s *grpc.Server, srv RouteMsgServer) {
	s.RegisterService(&_RouteMsg_serviceDesc, srv)
}

func _RouteMsg_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteMsgServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.RouteMsg/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteMsgServer).SendMsg(ctx, req.(*ComMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.RouteMsg",
	HandlerType: (*RouteMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _RouteMsg_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x6a, 0x84, 0x30,
	0x18, 0x84, 0x37, 0x5d, 0x36, 0xd6, 0xdf, 0x52, 0x6c, 0x0e, 0x45, 0x7a, 0x92, 0x40, 0x61, 0xa1,
	0x20, 0x25, 0x3d, 0xf4, 0x01, 0xc4, 0x63, 0x29, 0xfc, 0xa1, 0x0f, 0x60, 0x35, 0xe4, 0x14, 0x23,
	0xfe, 0xf1, 0x21, 0xfa, 0xd6, 0xc5, 0xa8, 0x20, 0xec, 0x6d, 0x86, 0xcc, 0x37, 0x4c, 0x7e, 0x48,
	0x1d, 0xd9, 0x6a, 0x9c, 0x7c, 0xf0, 0xe2, 0xdc, 0x79, 0x27, 0x15, 0xf0, 0xda, 0xbb, 0x2f, 0xb2,
	0x22, 0x87, 0xb3, 0x23, 0x5b, 0xb0, 0x92, 0x5d, 0x53, 0x5c, 0xa4, 0x78, 0x06, 0x4e, 0x66, 0xe8,
	0xcd, 0x54, 0xdc, 0x95, 0xec, 0x7a, 0xc1, 0xcd, 0xc9, 0x3f, 0x06, 0x0f, 0x3a, 0xb4, 0x61, 0x26,
	0x34, 0xa3, 0x9f, 0x82, 0x78, 0x07, 0x4e, 0xd1, 0x47, 0xfa, 0x51, 0x15, 0x55, 0xe7, 0x5d, 0x75,
	0x8c, 0xec, 0x66, 0xcb, 0x89, 0x57, 0x48, 0x9c, 0x21, 0x6a, 0xad, 0x89, 0xdd, 0x99, 0xca, 0x22,
	0xb2, 0x4e, 0xc1, 0xfd, 0x4d, 0x96, 0xc0, 0x57, 0x50, 0x64, 0x90, 0xe8, 0x9f, 0xba, 0x6e, 0xb4,
	0xce, 0x4f, 0x22, 0x85, 0x4b, 0x83, 0xf8, 0x8d, 0x39, 0x53, 0x9f, 0x70, 0x8f, 0x7e, 0x0e, 0x66,
	0xf9, 0xc1, 0x1b, 0x24, 0xda, 0x0c, 0xfd, 0x22, 0x8f, 0x75, 0x2f, 0x4f, 0x37, 0x73, 0xe4, 0xe9,
	0x97, 0xc7, 0x23, 0x7c, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xd5, 0x7a, 0xb4, 0x11, 0x01,
	0x00, 0x00,
}
