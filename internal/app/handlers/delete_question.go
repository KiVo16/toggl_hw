package handlers

import (
	"base/internal/constants"
	"base/internal/interfaces"
	"base/pkg/utils"
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
	userID := utils.GetIntFromContext(ctx, constants.ContextKeyUserID, 0)

	err := h.repo.DeleteQuestionByID(ctx, userID, req.ID)
	return nil, err
}
