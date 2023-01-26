package handlers

import (
	"base/internal/domain/model"
	"base/internal/interfaces"
	"base/pkg/pagination"
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
	list, err := h.repo.GetQuestions(ctx, req.Pages)
	return list, err
}
