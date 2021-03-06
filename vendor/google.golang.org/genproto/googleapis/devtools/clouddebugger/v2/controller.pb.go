// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto
// DO NOT EDIT!

/*
Package google_devtools_clouddebugger_v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/data.proto
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/debugger.proto

It has these top-level messages:
	RegisterDebuggeeRequest
	RegisterDebuggeeResponse
	ListActiveBreakpointsRequest
	ListActiveBreakpointsResponse
	UpdateActiveBreakpointRequest
	UpdateActiveBreakpointResponse
	FormatMessage
	StatusMessage
	SourceLocation
	Variable
	StackFrame
	Breakpoint
	Debuggee
	SetBreakpointRequest
	SetBreakpointResponse
	GetBreakpointRequest
	GetBreakpointResponse
	DeleteBreakpointRequest
	ListBreakpointsRequest
	ListBreakpointsResponse
	ListDebuggeesRequest
	ListDebuggeesResponse
*/
package google_devtools_clouddebugger_v2 // import "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "github.com/golang/protobuf/ptypes/empty"

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

// Request to register a debuggee.
type RegisterDebuggeeRequest struct {
	// Debuggee information to register.
	// The fields `project`, `uniquifier`, `description` and `agent_version`
	// of the debuggee must be set.
	Debuggee *Debuggee `protobuf:"bytes,1,opt,name=debuggee" json:"debuggee,omitempty"`
}

func (m *RegisterDebuggeeRequest) Reset()                    { *m = RegisterDebuggeeRequest{} }
func (m *RegisterDebuggeeRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterDebuggeeRequest) ProtoMessage()               {}
func (*RegisterDebuggeeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterDebuggeeRequest) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Response for registering a debuggee.
type RegisterDebuggeeResponse struct {
	// Debuggee resource.
	// The field `id` is guranteed to be set (in addition to the echoed fields).
	Debuggee *Debuggee `protobuf:"bytes,1,opt,name=debuggee" json:"debuggee,omitempty"`
}

func (m *RegisterDebuggeeResponse) Reset()                    { *m = RegisterDebuggeeResponse{} }
func (m *RegisterDebuggeeResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterDebuggeeResponse) ProtoMessage()               {}
func (*RegisterDebuggeeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterDebuggeeResponse) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Request to list active breakpoints.
type ListActiveBreakpointsRequest struct {
	// Identifies the debuggee.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// A wait token that, if specified, blocks the method call until the list
	// of active breakpoints has changed, or a server selected timeout has
	// expired.  The value should be set from the last returned response.
	WaitToken string `protobuf:"bytes,2,opt,name=wait_token,json=waitToken" json:"wait_token,omitempty"`
	// If set to `true`, returns `google.rpc.Code.OK` status and sets the
	// `wait_expired` response field to `true` when the server-selected timeout
	// has expired (recommended).
	//
	// If set to `false`, returns `google.rpc.Code.ABORTED` status when the
	// server-selected timeout has expired (deprecated).
	SuccessOnTimeout bool `protobuf:"varint,3,opt,name=success_on_timeout,json=successOnTimeout" json:"success_on_timeout,omitempty"`
}

func (m *ListActiveBreakpointsRequest) Reset()                    { *m = ListActiveBreakpointsRequest{} }
func (m *ListActiveBreakpointsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsRequest) ProtoMessage()               {}
func (*ListActiveBreakpointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response for listing active breakpoints.
type ListActiveBreakpointsResponse struct {
	// List of all active breakpoints.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints" json:"breakpoints,omitempty"`
	// A wait token that can be used in the next method call to block until
	// the list of breakpoints changes.
	NextWaitToken string `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken" json:"next_wait_token,omitempty"`
	// The `wait_expired` field is set to true by the server when the
	// request times out and the field `success_on_timeout` is set to true.
	WaitExpired bool `protobuf:"varint,3,opt,name=wait_expired,json=waitExpired" json:"wait_expired,omitempty"`
}

func (m *ListActiveBreakpointsResponse) Reset()                    { *m = ListActiveBreakpointsResponse{} }
func (m *ListActiveBreakpointsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsResponse) ProtoMessage()               {}
func (*ListActiveBreakpointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListActiveBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

// Request to update an active breakpoint.
type UpdateActiveBreakpointRequest struct {
	// Identifies the debuggee being debugged.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// Updated breakpoint information.
	// The field 'id' must be set.
	Breakpoint *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint" json:"breakpoint,omitempty"`
}

func (m *UpdateActiveBreakpointRequest) Reset()                    { *m = UpdateActiveBreakpointRequest{} }
func (m *UpdateActiveBreakpointRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointRequest) ProtoMessage()               {}
func (*UpdateActiveBreakpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateActiveBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Response for updating an active breakpoint.
// The message is defined to allow future extensions.
type UpdateActiveBreakpointResponse struct {
}

func (m *UpdateActiveBreakpointResponse) Reset()                    { *m = UpdateActiveBreakpointResponse{} }
func (m *UpdateActiveBreakpointResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointResponse) ProtoMessage()               {}
func (*UpdateActiveBreakpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*RegisterDebuggeeRequest)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeRequest")
	proto.RegisterType((*RegisterDebuggeeResponse)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeResponse")
	proto.RegisterType((*ListActiveBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsRequest")
	proto.RegisterType((*ListActiveBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsResponse")
	proto.RegisterType((*UpdateActiveBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointRequest")
	proto.RegisterType((*UpdateActiveBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Controller2 service

type Controller2Client interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application should call this method with
	// the same request content to get back the same stable `debuggee_id`. Agents
	// should call this method again whenever `google.rpc.Code.NOT_FOUND` is
	// returned from any controller method.
	//
	// This allows the controller service to disable the agent or recover from any
	// data loss. If the debuggee is disabled by the server, the response will
	// have `is_disabled` set to `true`.
	RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (location, condition, and expression
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller
	// service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expression` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error)
}

type controller2Client struct {
	cc *grpc.ClientConn
}

func NewController2Client(cc *grpc.ClientConn) Controller2Client {
	return &controller2Client{cc}
}

func (c *controller2Client) RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error) {
	out := new(RegisterDebuggeeResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error) {
	out := new(ListActiveBreakpointsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error) {
	out := new(UpdateActiveBreakpointResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Controller2 service

type Controller2Server interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application should call this method with
	// the same request content to get back the same stable `debuggee_id`. Agents
	// should call this method again whenever `google.rpc.Code.NOT_FOUND` is
	// returned from any controller method.
	//
	// This allows the controller service to disable the agent or recover from any
	// data loss. If the debuggee is disabled by the server, the response will
	// have `is_disabled` set to `true`.
	RegisterDebuggee(context.Context, *RegisterDebuggeeRequest) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (location, condition, and expression
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(context.Context, *ListActiveBreakpointsRequest) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller
	// service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expression` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(context.Context, *UpdateActiveBreakpointRequest) (*UpdateActiveBreakpointResponse, error)
}

func RegisterController2Server(s *grpc.Server, srv Controller2Server) {
	s.RegisterService(&_Controller2_serviceDesc, srv)
}

func _Controller2_RegisterDebuggee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDebuggeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).RegisterDebuggee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).RegisterDebuggee(ctx, req.(*RegisterDebuggeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_ListActiveBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, req.(*ListActiveBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_UpdateActiveBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActiveBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, req.(*UpdateActiveBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Controller2",
	HandlerType: (*Controller2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDebuggee",
			Handler:    _Controller2_RegisterDebuggee_Handler,
		},
		{
			MethodName: "ListActiveBreakpoints",
			Handler:    _Controller2_ListActiveBreakpoints_Handler,
		},
		{
			MethodName: "UpdateActiveBreakpoint",
			Handler:    _Controller2_UpdateActiveBreakpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0x66, 0x5a, 0x94, 0xf6, 0xac, 0xd2, 0x32, 0xa0, 0x86, 0xc5, 0x6a, 0xba, 0x48, 0x29, 0x25,
	0xec, 0xc0, 0xd6, 0x97, 0xe6, 0xc1, 0x4b, 0xbc, 0x51, 0xa9, 0x1a, 0x42, 0xc5, 0xc7, 0xb0, 0x97,
	0xd3, 0x71, 0xe8, 0x66, 0x66, 0xdd, 0x99, 0x8d, 0x2d, 0xa5, 0x2f, 0xbe, 0x2a, 0xbe, 0xf8, 0x63,
	0x04, 0xff, 0x86, 0xfe, 0x04, 0x9f, 0xfc, 0x15, 0x92, 0xdd, 0xcd, 0xa5, 0x97, 0x34, 0x69, 0xf4,
	0x25, 0x90, 0xef, 0xcc, 0xf7, 0x9d, 0xef, 0x3b, 0x7b, 0x66, 0xa0, 0xc9, 0x95, 0xe2, 0x31, 0xba,
	0x5c, 0xc5, 0xbe, 0xe4, 0xae, 0x4a, 0x39, 0xe3, 0x28, 0x93, 0x54, 0x19, 0xc5, 0x8a, 0x92, 0x9f,
	0x08, 0xcd, 0x22, 0xec, 0x1a, 0xa5, 0x62, 0xcd, 0xc2, 0x58, 0x65, 0x51, 0x84, 0x41, 0xc6, 0x39,
	0xa6, 0xac, 0xeb, 0xb1, 0x50, 0x49, 0x93, 0xaa, 0x38, 0xc6, 0xd4, 0xcd, 0x59, 0xb4, 0x5a, 0x2a,
	0xf6, 0x29, 0xee, 0x09, 0x8a, 0xdb, 0xf5, 0xec, 0xed, 0xe9, 0x7a, 0xfa, 0x89, 0x60, 0x1a, 0xd3,
	0xae, 0x08, 0x31, 0x54, 0x72, 0x4f, 0x70, 0xe6, 0x4b, 0xa9, 0x8c, 0x6f, 0x84, 0x92, 0xba, 0x68,
	0x66, 0xbf, 0xfc, 0x57, 0xfb, 0x91, 0x6f, 0xfc, 0x52, 0x6b, 0x93, 0x0b, 0xf3, 0x3e, 0x0b, 0xdc,
	0x50, 0x75, 0x58, 0xa1, 0xc7, 0xf2, 0x42, 0x90, 0xed, 0xb1, 0xc4, 0x1c, 0x26, 0xa8, 0x19, 0x76,
	0x12, 0x73, 0x58, 0xfc, 0x16, 0x24, 0xc7, 0x87, 0x5b, 0x2d, 0xe4, 0x42, 0x1b, 0x4c, 0x9f, 0x16,
	0xb2, 0xd8, 0xc2, 0x0f, 0x19, 0x6a, 0x43, 0x9f, 0xc3, 0x42, 0xd9, 0x09, 0x2b, 0xa4, 0x4a, 0xd6,
	0x2d, 0x6f, 0xc3, 0x9d, 0x34, 0x1b, 0x77, 0x20, 0x32, 0xe0, 0x3a, 0x01, 0x54, 0xce, 0xb6, 0xd0,
	0x89, 0x92, 0x1a, 0xff, 0x5b, 0x8f, 0x2f, 0x04, 0x6e, 0xef, 0x08, 0x6d, 0x1e, 0x87, 0x46, 0x74,
	0xb1, 0x91, 0xa2, 0xbf, 0x9f, 0x28, 0x21, 0x8d, 0xee, 0x87, 0xb9, 0x0b, 0x56, 0xff, 0x70, 0x5b,
	0x44, 0x79, 0xaf, 0xc5, 0x16, 0xf4, 0xa1, 0xed, 0x88, 0xae, 0x00, 0x7c, 0xf4, 0x85, 0x69, 0x1b,
	0xb5, 0x8f, 0xb2, 0x32, 0x97, 0xd7, 0x17, 0x7b, 0xc8, 0x6e, 0x0f, 0xa0, 0x35, 0xa0, 0x3a, 0x0b,
	0x43, 0xd4, 0xba, 0xad, 0x64, 0xdb, 0x88, 0x0e, 0xaa, 0xcc, 0x54, 0xe6, 0xab, 0x64, 0x7d, 0xa1,
	0xb5, 0x5c, 0x56, 0xde, 0xc8, 0xdd, 0x02, 0x77, 0x7e, 0x10, 0x58, 0x19, 0x63, 0xa7, 0x0c, 0xfe,
	0x1a, 0xac, 0x60, 0x08, 0x57, 0x48, 0x75, 0x7e, 0xdd, 0xf2, 0x6a, 0x93, 0xb3, 0x0f, 0xb5, 0x5a,
	0xa3, 0x02, 0x74, 0x0d, 0x96, 0x24, 0x1e, 0x98, 0xf6, 0x99, 0x0c, 0xd7, 0x7b, 0xf0, 0xbb, 0x41,
	0x8e, 0x55, 0xb8, 0x96, 0x1f, 0xc1, 0x83, 0x44, 0xa4, 0x18, 0x95, 0x09, 0xac, 0x1e, 0xf6, 0xac,
	0x80, 0x9c, 0xaf, 0x04, 0x56, 0xde, 0x26, 0x91, 0x6f, 0xf0, 0xb4, 0xfd, 0xa9, 0x87, 0xb9, 0x03,
	0x30, 0x34, 0x97, 0x1b, 0xb9, 0x6c, 0xb8, 0x11, 0xbe, 0x53, 0x85, 0x3b, 0xe3, 0xfc, 0x14, 0xd3,
	0xf4, 0x3e, 0x5f, 0x01, 0xeb, 0xc9, 0xe0, 0x22, 0x7b, 0xf4, 0x3b, 0x81, 0xe5, 0xd3, 0x3b, 0x47,
	0xb7, 0x26, 0x1b, 0x18, 0x73, 0x15, 0xec, 0xfa, 0x2c, 0xd4, 0xc2, 0x9b, 0x53, 0xfb, 0xf4, 0xf3,
	0xf7, 0xb7, 0xb9, 0x35, 0x67, 0xf5, 0xe4, 0x6b, 0xc3, 0xfa, 0xe3, 0xd2, 0x2c, 0x2d, 0xa9, 0x75,
	0xb2, 0x41, 0x7f, 0x11, 0xb8, 0x71, 0xee, 0xe6, 0xd0, 0x07, 0x93, 0x3d, 0x5c, 0x74, 0x03, 0xec,
	0x87, 0x33, 0xf3, 0xcb, 0x20, 0xf5, 0x3c, 0xc8, 0x7d, 0xea, 0x8d, 0x0d, 0x72, 0x34, 0xb2, 0x15,
	0xc7, 0x6c, 0x74, 0x3d, 0xff, 0x10, 0xb8, 0x79, 0xfe, 0x37, 0xa4, 0x53, 0xf8, 0xba, 0x70, 0x1b,
	0xed, 0x47, 0xb3, 0x0b, 0x94, 0xc9, 0x5e, 0xe5, 0xc9, 0x5e, 0xd8, 0x8d, 0xcb, 0x27, 0x63, 0x47,
	0xc3, 0x3f, 0xae, 0x88, 0x8e, 0xeb, 0x64, 0xa3, 0xb1, 0x05, 0xf7, 0x42, 0xd5, 0x99, 0xe8, 0xaa,
	0xb1, 0x34, 0x5c, 0xd9, 0x66, 0xef, 0x31, 0x6e, 0x92, 0xe0, 0x6a, 0xfe, 0x2a, 0x6f, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x89, 0xdc, 0x57, 0xe9, 0xd7, 0x06, 0x00, 0x00,
}
