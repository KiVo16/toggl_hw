package handlers

import (
	"base/internal/constants"
	"base/internal/domain/model"
	"base/internal/interfaces"
	"base/pkg/pagination"
	"base/pkg/utils"
	"context"
)

type GetQuestionsRequest struct {
	Pages pagination.Pagination
}

type GetQuestionsHandler Handler[GetQuestionsRequest, model.QuestionList]

type getQuestionsHandler struct {
	repo interfaces.Repository
}

func NewGetQuestionsHandler(repo interfaces.Repository) GetQuestionsHandler {
	if repo == nil {
		panic("NewGetQuestionsHandler: repo is nil")
	}

	return getQuestionsHandler{
		repo: repo,
	}
}

func (h getQuestionsHandler) Handle(ctx context.Context, req GetQuestionsRequest) (model.QuestionList, error) {
	req.Pages.DefaultIfNotSet()
	userID := utils.GetIntFromContext(ctx, constants.ContextKeyUserID, 0)

	list, err := h.repo.GetQuestions(ctx, userID, req.Pages)
	return list, err
}
