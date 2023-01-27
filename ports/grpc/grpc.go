package ports

import (
	"base/internal/app"
	"base/internal/app/handlers"
	"base/internal/domain/model"
	e "base/pkg/grpc/errors"
	"base/pkg/pagination"
	"base/pkg/utils"
	pb "base/ports/grpc/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.App
}

func NewGrpcServer(app app.App) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (s GrpcServer) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.CreateQuestionResponse, error) {
	question := req.Question
	if question == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: missing question value")
	}

	q := NewQuestionToModel(question)
	fQuestion, err := s.app.Handlers.CreateQuestion.Handle(ctx, handlers.CreateQuestionRequest{
		Question: q,
	})
	if err != nil {
		return nil, e.NewGRPCError(err).Err()
	}

	return &pb.CreateQuestionResponse{
		Question: NewQuestionFromModel(*fQuestion),
	}, nil
}

func (s GrpcServer) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*empty.Empty, error) {
	_, err := s.app.Handlers.DeleteQuestion.Handle(ctx, handlers.DeleteQuestionRequest{
		ID: int(req.Id),
	})
	if err != nil {
		return nil, e.NewGRPCError(err).Err()
	}

	return &empty.Empty{}, nil
}

func (s GrpcServer) GetQuestions(ctx context.Context, req *pb.GetQuestionsRequest) (*pb.GetQuestionsResponse, error) {
	r := handlers.GetQuestionsRequest{}

	pages := req.Pagination
	if pages != nil {
		r.Pages = pagination.Pagination{
			PageSize: int(pages.PageSize),
			Page:     int(pages.Page),
		}
	}

	questions, err := s.app.Handlers.GetQuestions.Handle(ctx, r)

	if err != nil {
		return nil, e.NewGRPCError(err).Err()
	}

	fQuestions := make([]*pb.Question, len(questions))
	for i, q := range questions {
		fQuestions[i] = NewQuestionFromModel(q)
	}

	return &pb.GetQuestionsResponse{
		Questions: fQuestions,
	}, nil
}

func (s GrpcServer) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.UpdateQuestionResponse, error) {
	ref := req.Ref
	if ref == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: missing ref value")
	}

	r := handlers.UpdateQuestionRequest{
		ID: int(req.Id),
	}

	if utils.Contains(req.PropsToUpdate, pb.UpdateQuestionRequest_BODY) {
		r.Body = &ref.Body
	}

	if utils.Contains(req.PropsToUpdate, pb.UpdateQuestionRequest_OPTIONS) {
		options := make([]model.Option, len(ref.Options))
		for i, o := range ref.Options {
			options[i] = NewOptionToModel(o)
		}

		r.Options = &options
	}

	fQuestion, err := s.app.Handlers.UpdateQuestion.Handle(ctx, r)
	if err != nil {
		return nil, e.NewGRPCError(err).Err()
	}

	return &pb.UpdateQuestionResponse{
		Question: NewQuestionFromModel(*fQuestion),
	}, nil
}
