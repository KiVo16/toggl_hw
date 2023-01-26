package interfaces

import (
	"base/internal/domain/model"
	"base/pkg/pagination"
	"context"
)

type Repository interface {
	CreateQuestion(ctx context.Context, question model.Question) error
	DeleteQuestionByID(ctx context.Context, id int) error
	GetQuestions(ctx context.Context, pagination pagination.Pagination) (model.QuestionList, error)
	UpdateQuestion(
		ctx context.Context,
		id int,
		updateFn func(q *model.Question) (*model.Question, error),
	) error
}
