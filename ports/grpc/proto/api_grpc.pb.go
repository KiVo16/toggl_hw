// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api.proto

package __toggl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QuestionsServiceClient is the client API for QuestionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionsServiceClient interface {
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error)
	DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (*GetQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*UpdateQuestionResponse, error)
}

type questionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionsServiceClient(cc grpc.ClientConnInterface) QuestionsServiceClient {
	return &questionsServiceClient{cc}
}

func (c *questionsServiceClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error) {
	out := new(CreateQuestionResponse)
	err := c.cc.Invoke(ctx, "/toggl.QuestionsService/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionsServiceClient) DeleteQuestion(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/toggl.QuestionsService/DeleteQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionsServiceClient) GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (*GetQuestionsResponse, error) {
	out := new(GetQuestionsResponse)
	err := c.cc.Invoke(ctx, "/toggl.QuestionsService/GetQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionsServiceClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*UpdateQuestionResponse, error) {
	out := new(UpdateQuestionResponse)
	err := c.cc.Invoke(ctx, "/toggl.QuestionsService/UpdateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionsServiceServer is the server API for QuestionsService service.
// All implementations should embed UnimplementedQuestionsServiceServer
// for forward compatibility
type QuestionsServiceServer interface {
	CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error)
	DeleteQuestion(context.Context, *DeleteQuestionRequest) (*emptypb.Empty, error)
	GetQuestions(context.Context, *GetQuestionsRequest) (*GetQuestionsResponse, error)
	UpdateQuestion(context.Context, *UpdateQuestionRequest) (*UpdateQuestionResponse, error)
}

// UnimplementedQuestionsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQuestionsServiceServer struct {
}

func (UnimplementedQuestionsServiceServer) CreateQuestion(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedQuestionsServiceServer) DeleteQuestion(context.Context, *DeleteQuestionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (UnimplementedQuestionsServiceServer) GetQuestions(context.Context, *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestions not implemented")
}
func (UnimplementedQuestionsServiceServer) UpdateQuestion(context.Context, *UpdateQuestionRequest) (*UpdateQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}

// UnsafeQuestionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionsServiceServer will
// result in compilation errors.
type UnsafeQuestionsServiceServer interface {
	mustEmbedUnimplementedQuestionsServiceServer()
}

func RegisterQuestionsServiceServer(s grpc.ServiceRegistrar, srv QuestionsServiceServer) {
	s.RegisterService(&QuestionsService_ServiceDesc, srv)
}

func _QuestionsService_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toggl.QuestionsService/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionsService_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toggl.QuestionsService/DeleteQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).DeleteQuestion(ctx, req.(*DeleteQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionsService_GetQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).GetQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toggl.QuestionsService/GetQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).GetQuestions(ctx, req.(*GetQuestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionsService_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toggl.QuestionsService/UpdateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).UpdateQuestion(ctx, req.(*UpdateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionsService_ServiceDesc is the grpc.ServiceDesc for QuestionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "toggl.QuestionsService",
	HandlerType: (*QuestionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _QuestionsService_CreateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _QuestionsService_DeleteQuestion_Handler,
		},
		{
			MethodName: "GetQuestions",
			Handler:    _QuestionsService_GetQuestions_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _QuestionsService_UpdateQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}