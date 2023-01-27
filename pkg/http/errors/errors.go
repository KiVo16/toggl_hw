package ports

import (
	e "base/internal/errors"
	"encoding/json"
	"log"
	"net/http"
)

var errorsStatusCodesMap = map[error]int{
	e.ErrQuestionEmptyBody:            http.StatusBadRequest,
	e.ErrQuestionMissingOptions:       http.StatusBadRequest,
	e.ErrQuestionMissingCorrectOption: http.StatusBadRequest,
	e.ErrOptionEmptyBody:              http.StatusBadRequest,
	e.ErrQuestionUserIDZero:           http.StatusBadRequest,
	e.ErrQuestionNotFound:             http.StatusNotFound,
}

type HttpError struct {
	err  error
	code *int

	ErrorMsg string `json:"error"`
}

func NewHttpError(err error) *HttpError {
	return &HttpError{
		err:      err,
		ErrorMsg: err.Error(),
	}
}

func (e *HttpError) WithCode(code int) *HttpError {
	e.code = &code
	return e
}

func (e HttpError) Handle(w http.ResponseWriter) {
	finalCode, ok := errorsStatusCodesMap[e.err]
	if !ok {
		if e.code != nil {
			finalCode = *e.code
		} else {
			finalCode = http.StatusInternalServerError
		}
	}

	j, err := json.Marshal(e)
	if err != nil {
		log.Printf("HttpError: marshal error = %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		log.Printf("HttpError: write error = %v", err)
	}

	w.WriteHeader(finalCode)
}
