package handlers

import (
	"base/internal/domain/model"
	"base/internal/interfaces"
	"context"
)

type CreateQuestionRequest struct {
	Question model.Question
}

type CreateQuestionHandler Handler[CreateQuestionRequest, interface{}]

type createRequestHandler struct {
	repo interfaces.Repository
}

func NewCreateQuestionHandler(repo interfaces.Repository) CreateQuestionHandler {
	if repo == nil {
		panic("NewCreateQuestionHandler: repo is nil")
	}

	return createRequestHandler{
		repo: repo,
	}
}

func (h createRequestHandler) Handle(ctx context.Context, req CreateQuestionRequest) (interface{}, error) {
	err := req.Question.Validate()
	if err != nil {
		return nil, err
	}

	err = h.repo.CreateQuestion(ctx, req.Question)
	return nil, err
}
