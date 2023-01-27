package ports

import (
	e "base/internal/errors"
	"net/http"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

var errorsStatusCodesMap = map[error]codes.Code{
	e.ErrQuestionEmptyBody:      codes.InvalidArgument,
	e.ErrQuestionMissingOptions: codes.InvalidArgument,
	e.ErrOptionEmptyBody:        codes.InvalidArgument,
	e.ErrQuestionUserIDZero:     codes.InvalidArgument,
	e.ErrQuestionNotFound:       codes.NotFound,
}

type GRPCError struct {
	err  error
	code *codes.Code

	ErrorMsg string `json:"error"`
}

func NewGRPCError(err error) *GRPCError {
	return &GRPCError{
		err:      err,
		ErrorMsg: err.Error(),
	}
}

func (e *GRPCError) WithCode(code codes.Code) *GRPCError {
	e.code = &code
	return e
}

func (e GRPCError) Err() error {
	finalCode, ok := errorsStatusCodesMap[e.err]
	if !ok {
		if e.code != nil {
			finalCode = *e.code
		} else {
			finalCode = http.StatusInternalServerError
		}
	}

	return status.Error(finalCode, e.ErrorMsg)
}
