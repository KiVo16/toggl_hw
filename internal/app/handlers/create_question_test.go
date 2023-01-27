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

func TestCreateQuestion_Handle(t *testing.T) {

	testUserId := 10

	tests := []struct {
		name       string
		q          model.Question
		beforeTest func(r *interfaces.MockRepository) error
		err        error
	}{
		{
			name: "create question without body",
			q: model.Question{
				ID:     1,
				UserID: testUserId,
				Options: []model.Option{
					model.Option{
						Body:    "test",
						Correct: true,
					},
					model.Option{
						Body:    "test2",
						Correct: false,
					},
				},
			},
			beforeTest: func(r *interfaces.MockRepository) error {
				return nil
			},
			err: errors.ErrQuestionEmptyBody,
		},
		{
			name: "create question without single options",
			q: model.Question{
				ID:     1,
				UserID: testUserId,
				Body:   "test",
			},
			beforeTest: func(r *interfaces.MockRepository) error {
				return nil
			},
			err: errors.ErrQuestionMissingOptions,
		},
		{
			name: "create question without correct option",
			q: model.Question{
				ID:     1,
				UserID: testUserId,
				Body:   "test",
				Options: []model.Option{
					model.Option{
						Body:    "test",
						Correct: false,
					},
					model.Option{
						Body:    "test2",
						Correct: false,
					},
				},
			},
			beforeTest: func(r *interfaces.MockRepository) error {
				return nil
			},
			err: errors.ErrQuestionMissingCorrectOption,
		},
		{
			name: "create question",
			q: model.Question{
				ID:     1,
				UserID: testUserId,
				Body:   "test",
				Options: []model.Option{
					model.Option{
						Body:    "test",
						Correct: true,
					},
					model.Option{
						Body:    "test2",
						Correct: false,
					},
				},
			},
			beforeTest: func(r *interfaces.MockRepository) error {
				return nil
			},
			err: nil,
		},
	}

	for _, tt := range tests {

		extraMsg := fmt.Sprintf("Test name: %s", tt.name)

		repo := interfaces.NewMockRepository()
		err := tt.beforeTest(repo)
		assert.Nil(t, err, extraMsg)

		ctx := context.WithValue(context.Background(), constants.ContextKeyUserID, testUserId)
		req := CreateQuestionRequest{
			Question: tt.q,
		}

		handler := NewCreateQuestionHandler(repo)
		_, err = handler.Handle(ctx, req)
		assert.Equal(t, tt.err, err, extraMsg)
	}
}
