package model

import (
	"errors"
	"strings"
)

var (
	ErrQuestionMissingOptions = errors.New("errQuestionMissingOptions")
	ErrQuestionEmptyBody      = errors.New("errQuestionEmptyBody")
)

type Question struct {
	Id      int
	Body    string
	Options []Option
}

type QuestionList = []Question

func (q Question) Validate() error {

	// it's just for an example - it probably should accept generating question without options
	if len(q.Options) == 0 {
		return ErrQuestionMissingOptions
	}

	if len(strings.TrimSpace(q.Body)) == 0 {
		return ErrQuestionEmptyBody
	}

	for _, o := range q.Options {
		err := o.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
