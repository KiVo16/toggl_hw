package handlers

import (
	"context"

	"base/internal/interfaces"
)

type Handler[T any, I any] interface {
	Handle(ctx context.Context, data T) (I, error)
}

type Handlers struct {
	CreateQuestion CreateQuestionHandler
	DeleteQuestion DeleteQuestionHandler
	GetQuestions   GetQuestionsHandler
	UpdateQuestion UpdateQuestionHandler
}

func NewHandlers(repo interfaces.Repository) Handlers {

	return Handlers{
		CreateQuestion: NewCreateQuestionHandler(repo),
		DeleteQuestion: NewDeleteQuestionHandler(repo),
		GetQuestions:   NewGetQuestionsHandler(repo),
		UpdateQuestion: NewUpdateQuestionHandler(repo),
	}
}
