// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: google/bytestream/bytestream.proto

package bytestream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ByteStreamClient is the client API for ByteStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ByteStreamClient interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (ByteStream_ReadClient, error)
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(ctx context.Context, opts ...grpc.CallOption) (ByteStream_WriteClient, error)
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(ctx context.Context, in *QueryWriteStatusRequest, opts ...grpc.CallOption) (*QueryWriteStatusResponse, error)
}

type byteStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewByteStreamClient(cc grpc.ClientConnInterface) ByteStreamClient {
	return &byteStreamClient{cc}
}

func (c *byteStreamClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (ByteStream_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ByteStream_ServiceDesc.Streams[0], "/google.bytestream.ByteStream/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &byteStreamReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ByteStream_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type byteStreamReadClient struct {
	grpc.ClientStream
}

func (x *byteStreamReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *byteStreamClient) Write(ctx context.Context, opts ...grpc.CallOption) (ByteStream_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ByteStream_ServiceDesc.Streams[1], "/google.bytestream.ByteStream/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &byteStreamWriteClient{stream}
	return x, nil
}

type ByteStream_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type byteStreamWriteClient struct {
	grpc.ClientStream
}

func (x *byteStreamWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *byteStreamWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *byteStreamClient) QueryWriteStatus(ctx context.Context, in *QueryWriteStatusRequest, opts ...grpc.CallOption) (*QueryWriteStatusResponse, error) {
	out := new(QueryWriteStatusResponse)
	err := c.cc.Invoke(ctx, "/google.bytestream.ByteStream/QueryWriteStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ByteStreamServer is the server API for ByteStream service.
// All implementations must embed UnimplementedByteStreamServer
// for forward compatibility
type ByteStreamServer interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(*ReadRequest, ByteStream_ReadServer) error
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(ByteStream_WriteServer) error
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(context.Context, *QueryWriteStatusRequest) (*QueryWriteStatusResponse, error)
	mustEmbedUnimplementedByteStreamServer()
}

// UnimplementedByteStreamServer must be embedded to have forward compatible implementations.
type UnimplementedByteStreamServer struct {
}

func (UnimplementedByteStreamServer) Read(*ReadRequest, ByteStream_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedByteStreamServer) Write(ByteStream_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedByteStreamServer) QueryWriteStatus(context.Context, *QueryWriteStatusRequest) (*QueryWriteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWriteStatus not implemented")
}
func (UnimplementedByteStreamServer) mustEmbedUnimplementedByteStreamServer() {}

// UnsafeByteStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ByteStreamServer will
// result in compilation errors.
type UnsafeByteStreamServer interface {
	mustEmbedUnimplementedByteStreamServer()
}

func RegisterByteStreamServer(s grpc.ServiceRegistrar, srv ByteStreamServer) {
	s.RegisterService(&ByteStream_ServiceDesc, srv)
}

func _ByteStream_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ByteStreamServer).Read(m, &byteStreamReadServer{stream})
}

type ByteStream_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type byteStreamReadServer struct {
	grpc.ServerStream
}

func (x *byteStreamReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ByteStream_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ByteStreamServer).Write(&byteStreamWriteServer{stream})
}

type ByteStream_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type byteStreamWriteServer struct {
	grpc.ServerStream
}

func (x *byteStreamWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *byteStreamWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ByteStream_QueryWriteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWriteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ByteStreamServer).QueryWriteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bytestream.ByteStream/QueryWriteStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ByteStreamServer).QueryWriteStatus(ctx, req.(*QueryWriteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ByteStream_ServiceDesc is the grpc.ServiceDesc for ByteStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ByteStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.bytestream.ByteStream",
	HandlerType: (*ByteStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryWriteStatus",
			Handler:    _ByteStream_QueryWriteStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _ByteStream_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Write",
			Handler:       _ByteStream_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "google/bytestream/bytestream.proto",
}