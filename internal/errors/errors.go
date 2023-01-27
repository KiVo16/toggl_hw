package errors

import "errors"

var (
	ErrOptionEmptyBody              = errors.New("option's body cannot be empty")
	ErrQuestionMissingOptions       = errors.New("question must have at least 2 options")
	ErrQuestionMissingCorrectOption = errors.New("at least 1 option must be marked as correct")
	ErrQuestionEmptyBody            = errors.New("question's body cannot be empty")
	ErrQuestionUserIDZero           = errors.New("question's userID cannot be equal 0")

	ErrQuestionNotFound = errors.New("question not found")
	ErrInternal         = errors.New("internal error")
)
