package db

import (
	"base/internal/domain/model"

	"github.com/uptrace/bun"
)

type Question struct {
	bun.BaseModel `bun:"table:questions,alias:q"`

	ID      int64     `bun:"id,pk,autoincrement"`
	UserID  int64     `bun:"user_id"`
	Body    string    `bun:"body"`
	Options []*Option `bun:"rel:has-many,join:id=question_id"`
}

func (q Question) ToModel() model.Question {
	question := model.Question{
		ID:     int(q.ID),
		UserID: int(q.UserID),
		Body:   q.Body,
	}

	for _, o := range q.Options {
		question.Options = append(question.Options, o.ToModel())
	}

	return question
}

type Option struct {
	bun.BaseModel `bun:"table:questions_options,alias:qo"`

	ID         int64  `bun:"id,pk,autoincrement"`
	QuestionID int64  `bun:"question_id"`
	Body       string `bun:"body"`
	Correct    bool   `bun:"correct"`
}

func (o Option) ToModel() model.Option {
	return model.Option{
		ID:      int(o.ID),
		Body:    o.Body,
		Correct: o.Correct,
	}
}

func NewQuestion(q model.Question) *Question {
	question := Question{
		ID:     int64(q.ID),
		UserID: int64(q.UserID),
		Body:   q.Body,
	}

	for _, o := range q.Options {
		option := Option{
			ID:      int64(o.ID),
			Body:    o.Body,
			Correct: o.Correct,
		}
		question.Options = append(question.Options, &option)
	}

	return &question
}

func (q *Question) SyncOptionsQuestionIDs() *Question {
	for i, o := range q.Options {
		if o == nil {
			continue
		}

		q.Options[i].QuestionID = q.ID
	}
	return q
}
