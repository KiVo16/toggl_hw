package handlers

import (
	"base/internal/constants"
	"base/internal/domain/model"
	"base/internal/errors"
	"base/internal/interfaces"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteQuestion_Handle(t *testing.T) {

	testUserId := 10

	tests := []struct {
		name       string
		id         int
		beforeTest func(r *interfaces.MockRepository) error
		err        error
	}{
		{
			name: "delete existing question",
			id:   1,
			beforeTest: func(r *interfaces.MockRepository) error {
				q := model.Question{
					ID:     1,
					UserID: testUserId,
				}
				_, err := r.CreateQuestion(context.Background(), q)
				return err
			},
			err: nil,
		},
		{
			name: "delete question that don't exists",
			id:   1,
			beforeTest: func(r *interfaces.MockRepository) error {
				return nil
			},
			err: errors.ErrQuestionNotFound,
		},
	}

	for _, tt := range tests {

		extraMsg := fmt.Sprintf("Test name: %s", tt.name)

		repo := interfaces.NewMockRepository()
		err := tt.beforeTest(repo)
		assert.Nil(t, err, extraMsg)

		ctx := context.WithValue(context.Background(), constants.ContextKeyUserID, testUserId)
		req := DeleteQuestionRequest{
			ID: tt.id,
		}

		handler := NewDeleteQuestionHandler(repo)
		_, err = handler.Handle(ctx, req)
		assert.Equal(t, tt.err, err, extraMsg)
	}
}
