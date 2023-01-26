package handlers

import (
	"base/internal/interfaces"
	"context"
)

type UpdateQuestionRequest struct {
	ID int
}

type UpdateQuestionHandler Handler[UpdateQuestionRequest, interface{}]

type updateQuestionHandler struct {
	repo interfaces.Repository
}

func NewUpdateQuestionHandler(repo interfaces.Repository) UpdateQuestionHandler {
	if repo == nil {
		panic("NewUpdateQuestionHandler: repo is nil")
	}

	return updateQuestionHandler{
		repo: repo,
	}
}

func (h updateQuestionHandler) Handle(ctx context.Context, req UpdateQuestionRequest) (interface{}, error) {
	err := h.repo.DeleteQuestionByID(ctx, req.ID)
	return nil, err
}
