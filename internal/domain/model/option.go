package model

import (
	"errors"
	"strings"
)

var (
	ErrOptionEmptyBody = errors.New("errOptionEmptyBody")
)

type Option struct {
	Id      int
	Body    string
	Correct bool
}

func (o Option) Validate() error {
	if len(strings.TrimSpace(o.Body)) == 0 {
		return ErrOptionEmptyBody
	}
	return nil
}
