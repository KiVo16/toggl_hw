package handlers

import (
	"base/internal/constants"
	"base/internal/domain/model"
	"base/internal/interfaces"
	"base/pkg/utils"
	"context"
)

type CreateQuestionRequest struct {
	Question model.Question
}

type CreateQuestionHandler Handler[CreateQuestionRequest, *model.Question]

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

func (h createRequestHandler) Handle(ctx context.Context, req CreateQuestionRequest) (*model.Question, error) {
	q := req.Question
	q.UserID = utils.GetIntFromContext(ctx, constants.ContextKeyUserID, 0)

	err := q.Validate()
	if err != nil {
		return nil, err
	}

	question, err := h.repo.CreateQuestion(ctx, q)
	return question, err
}
