// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file_service.proto

package devtools_goma

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	api "go.chromium.org/goma/server/proto/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("file/file_service.proto", fileDescriptor_f713fbe7eb11f137) }

var fileDescriptor_f713fbe7eb11f137 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcb, 0xcc, 0x49,
	0xd5, 0x07, 0x11, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x29, 0xa9, 0x65, 0x25, 0xf9, 0xf9, 0x39, 0xc5, 0xf1, 0xe9, 0xf9, 0xb9, 0x89,
	0x52, 0xc2, 0x89, 0x05, 0x99, 0xfa, 0x20, 0x56, 0x7c, 0x4a, 0x62, 0x49, 0x22, 0x44, 0x8d, 0xd1,
	0x12, 0x46, 0x2e, 0x6e, 0xb7, 0xcc, 0x9c, 0xd4, 0x60, 0x88, 0x4e, 0x21, 0x0f, 0x2e, 0xce, 0xe0,
	0x92, 0xfc, 0xa2, 0x54, 0x90, 0x98, 0x90, 0xb4, 0x1e, 0x8a, 0x09, 0x7a, 0x70, 0x99, 0xa0, 0xd4,
	0x42, 0x29, 0x19, 0xdc, 0x92, 0xc5, 0x05, 0x4a, 0x0c, 0x42, 0xde, 0x5c, 0x5c, 0x3e, 0xf9, 0xf9,
	0xd9, 0xa5, 0x05, 0x60, 0xa3, 0xd0, 0x55, 0x23, 0xa4, 0x40, 0x66, 0xc9, 0xe2, 0x91, 0x05, 0x19,
	0xe6, 0xc4, 0xd9, 0xc0, 0xc8, 0xd0, 0xc1, 0xc8, 0x30, 0x81, 0x91, 0x01, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xab, 0xbd, 0x67, 0x76, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	StoreFile(ctx context.Context, in *api.StoreFileReq, opts ...grpc.CallOption) (*api.StoreFileResp, error)
	LookupFile(ctx context.Context, in *api.LookupFileReq, opts ...grpc.CallOption) (*api.LookupFileResp, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) StoreFile(ctx context.Context, in *api.StoreFileReq, opts ...grpc.CallOption) (*api.StoreFileResp, error) {
	out := new(api.StoreFileResp)
	err := c.cc.Invoke(ctx, "/devtools_goma.FileService/StoreFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) LookupFile(ctx context.Context, in *api.LookupFileReq, opts ...grpc.CallOption) (*api.LookupFileResp, error) {
	out := new(api.LookupFileResp)
	err := c.cc.Invoke(ctx, "/devtools_goma.FileService/LookupFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	StoreFile(context.Context, *api.StoreFileReq) (*api.StoreFileResp, error)
	LookupFile(context.Context, *api.LookupFileReq) (*api.LookupFileResp, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) StoreFile(ctx context.Context, req *api.StoreFileReq) (*api.StoreFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFile not implemented")
}
func (*UnimplementedFileServiceServer) LookupFile(ctx context.Context, req *api.LookupFileReq) (*api.LookupFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupFile not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_StoreFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.StoreFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).StoreFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devtools_goma.FileService/StoreFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).StoreFile(ctx, req.(*api.StoreFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_LookupFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.LookupFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).LookupFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devtools_goma.FileService/LookupFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).LookupFile(ctx, req.(*api.LookupFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "devtools_goma.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreFile",
			Handler:    _FileService_StoreFile_Handler,
		},
		{
			MethodName: "LookupFile",
			Handler:    _FileService_LookupFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/file_service.proto",
}
