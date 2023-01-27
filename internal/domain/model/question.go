package model

import (
	e "base/internal/errors"
	"strings"
)

type Question struct {
	ID      int
	UserID  int
	Body    string
	Options []Option
}

type QuestionList = []Question

func (q Question) Validate() error {

	// it's just for an example - it probably should accept generating question without options
	if len(q.Options) < 2 {
		return e.ErrQuestionMissingOptions
	}

	isOneCorrect := false
	for _, o := range q.Options {
		if o.Correct {
			isOneCorrect = true
			break
		}
	}

	if !isOneCorrect {
		return e.ErrQuestionMissingCorrectOption
	}

	if q.UserID == 0 {
		return e.ErrQuestionUserIDZero
	}

	if len(strings.TrimSpace(q.Body)) == 0 {
		return e.ErrQuestionEmptyBody
	}

	for _, o := range q.Options {
		err := o.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
