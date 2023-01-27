package handlers

import (
	"base/internal/constants"
	"base/internal/domain/model"
	"base/internal/interfaces"
	"base/pkg/utils"
	"context"
)

type UpdateQuestionRequest struct {
	ID      int
	Body    *string
	Options *[]model.Option
}

type UpdateQuestionHandler Handler[UpdateQuestionRequest, *model.Question]

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

func (h updateQuestionHandler) Handle(ctx context.Context, req UpdateQuestionRequest) (*model.Question, error) {
	userID := utils.GetIntFromContext(ctx, constants.ContextKeyUserID, 0)

	question, err := h.repo.UpdateQuestion(ctx, userID, req.ID, func(q *model.Question) (*model.Question, error) {
		if req.Body != nil {
			q.Body = *req.Body
		}

		if req.Options != nil {
			q.Options = *req.Options
		}

		err := q.Validate()
		if err != nil {
			return nil, err
		}

		return q, nil
	})
	return question, err
}
