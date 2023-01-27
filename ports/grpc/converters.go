package ports

import (
	"base/internal/domain/model"
	pb "base/ports/grpc/proto"
)

func NewOptionFromModel(o model.Option) *pb.Option {
	option := pb.Option{
		Id:      int64(o.ID),
		Body:    o.Body,
		Correct: o.Correct,
	}

	return &option
}

func NewQuestionFromModel(q model.Question) *pb.Question {
	question := pb.Question{
		Id:   int64(q.ID),
		Body: q.Body,
	}

	options := make([]*pb.Option, len(q.Options))
	for i, o := range q.Options {
		options[i] = NewOptionFromModel(o)
	}

	if len(options) > 0 {
		question.Options = options
	}

	return &question
}

func NewOptionToModel(o *pb.Option) model.Option {
	option := model.Option{
		Body:    o.Body,
		Correct: o.Correct,
	}

	if o.Id != 0 {
		option.ID = int(o.Id)
	}

	return option
}

func NewQuestionToModel(q *pb.Question) model.Question {
	question := model.Question{
		Body: q.Body,
	}

	if q.Id != 0 {
		question.ID = int(q.Id)
	}

	for _, o := range q.Options {
		if o == nil {
			continue
		}

		option := NewOptionToModel(o)
		question.Options = append(question.Options, option)
	}

	return question
}
