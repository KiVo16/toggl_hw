package handlers

import (
	"base/internal/interfaces"
	"context"
)

type DeleteQuestionRequest struct {
	ID int
}

type DeleteQuestionHandler Handler[DeleteQuestionRequest, interface{}]

type deleteQuestionHandler struct {
	repo interfaces.Repository
}

func NewDeleteQuestionHandler(repo interfaces.Repository) DeleteQuestionHandler {
	if repo == nil {
		panic("NewDeleteQuestionHandler: repo is nil")
	}

	return deleteQuestionHandler{
		repo: repo,
	}
}

func (h deleteQuestionHandler) Handle(ctx context.Context, req DeleteQuestionRequest) (interface{}, error) {
	err := h.repo.DeleteQuestionByID(ctx, req.ID)
	return nil, err
}
