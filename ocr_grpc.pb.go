// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocr_v1

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

// OCRServiceClient is the client API for OCRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCRServiceClient interface {
	// Perform a new ocr process request
	Recognize(ctx context.Context, in *OcrRequest, opts ...grpc.CallOption) (*OcrResponse, error)
	// Check the status of request if queued
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*OcrResponse, error)
}

type oCRServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOCRServiceClient(cc grpc.ClientConnInterface) OCRServiceClient {
	return &oCRServiceClient{cc}
}

func (c *oCRServiceClient) Recognize(ctx context.Context, in *OcrRequest, opts ...grpc.CallOption) (*OcrResponse, error) {
	out := new(OcrResponse)
	err := c.cc.Invoke(ctx, "/apis.OCRService/Recognize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCRServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*OcrResponse, error) {
	out := new(OcrResponse)
	err := c.cc.Invoke(ctx, "/apis.OCRService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCRServiceServer is the server API for OCRService service.
// All implementations must embed UnimplementedOCRServiceServer
// for forward compatibility
type OCRServiceServer interface {
	// Perform a new ocr process request
	Recognize(context.Context, *OcrRequest) (*OcrResponse, error)
	// Check the status of request if queued
	Status(context.Context, *StatusRequest) (*OcrResponse, error)
	mustEmbedUnimplementedOCRServiceServer()
}

// UnimplementedOCRServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOCRServiceServer struct {
}

func (UnimplementedOCRServiceServer) Recognize(context.Context, *OcrRequest) (*OcrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}
func (UnimplementedOCRServiceServer) Status(context.Context, *StatusRequest) (*OcrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedOCRServiceServer) mustEmbedUnimplementedOCRServiceServer() {}

// UnsafeOCRServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCRServiceServer will
// result in compilation errors.
type UnsafeOCRServiceServer interface {
	mustEmbedUnimplementedOCRServiceServer()
}

func RegisterOCRServiceServer(s grpc.ServiceRegistrar, srv OCRServiceServer) {
	s.RegisterService(&OCRService_ServiceDesc, srv)
}

func _OCRService_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OcrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServiceServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.OCRService/Recognize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServiceServer).Recognize(ctx, req.(*OcrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCRService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.OCRService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OCRService_ServiceDesc is the grpc.ServiceDesc for OCRService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCRService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.OCRService",
	HandlerType: (*OCRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _OCRService_Recognize_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _OCRService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/ocr.proto",
}
