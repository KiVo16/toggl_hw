package interfaces

import (
	"base/internal/domain/model"
	"base/pkg/pagination"
	"context"
)

type Repository interface {
	CreateQuestion(ctx context.Context, question model.Question) (*model.Question, error)
	DeleteQuestionByID(ctx context.Context, userId, id int) error
	GetQuestions(ctx context.Context, userId int, pagination pagination.Pagination) (model.QuestionList, error)
	UpdateQuestion(
		ctx context.Context,
		userId int,
		id int,
		updateFn func(q *model.Question) (*model.Question, error),
	) (*model.Question, error)
}
