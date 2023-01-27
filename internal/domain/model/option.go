package model

import (
	e "base/internal/errors"
	"strings"
)

type Option struct {
	ID      int
	Body    string
	Correct bool
}

func (o Option) Validate() error {
	if len(strings.TrimSpace(o.Body)) == 0 {
		return e.ErrOptionEmptyBody
	}
	return nil
}
