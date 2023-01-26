package ports

import (
	"base/internal/domain/model"
)

func NewOptionFromModel(o model.Option) Option {
	option := Option{
		Body:    o.Body,
		Correct: o.Correct,
	}

	if o.Id != 0 {
		option.Id = &o.Id
	}

	return option
}

func NewQuestionFromModel(q model.Question) Question {
	question := Question{
		Body: q.Body,
	}

	options := make([]Option, len(q.Options))
	for i, o := range q.Options {
		options[i] = NewOptionFromModel(o)
	}

	if len(options) > 0 {
		question.Options = &options
	}

	if q.Id != 0 {
		question.Id = &q.Id
	}

	return question
}

func NewOptionToModel(o Option) model.Option {
	option := model.Option{
		Body:    o.Body,
		Correct: o.Correct,
	}

	if o.Id != nil {
		option.Id = *o.Id
	}

	return option
}

func NewQuestionToModel(q Question) model.Question {
	question := model.Question{
		Body: q.Body,
	}

	if q.Id != nil {
		question.Id = *q.Id
	}

	if q.Options != nil {
		for _, o := range *q.Options {
			option := NewOptionToModel(o)
			question.Options = append(question.Options, option)
		}
	}

	return question
}
