package interfaces

import (
	"base/internal/domain/model"
	"base/internal/errors"
	"base/pkg/pagination"
	"base/pkg/utils"
	"context"
)

type MockRepository struct {
	data  map[int]model.Question
	order map[int][]int
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		data:  map[int]model.Question{},
		order: map[int][]int{},
	}
}

func (m MockRepository) getQuestion(id, userId int) (*model.Question, error) {
	q, ok := m.data[id]
	if !ok {
		return nil, errors.ErrQuestionNotFound
	}

	if q.UserID != userId {
		return nil, errors.ErrQuestionNotFound
	}

	return &q, nil
}

func (m *MockRepository) CreateQuestion(ctx context.Context, question model.Question) (*model.Question, error) {
	m.data[question.ID] = question

	v, ok := m.order[question.UserID]
	if !ok {
		m.order[question.UserID] = []int{question.ID}
		return &question, nil
	}

	v = append(v, question.ID)
	m.order[question.UserID] = v
	return &question, nil
}

func (m *MockRepository) DeleteQuestionByID(ctx context.Context, userId, id int) error {
	q, err := m.getQuestion(id, userId)
	if err != nil {
		return err
	}

	delete(m.data, q.ID)

	ov, ok := m.order[q.UserID]
	if ok {
		for i, v := range ov {
			if v != q.ID {
				continue
			}

			ov = utils.DeleteAtIndex(ov, i)
		}
		m.order[q.UserID] = ov
	}

	return nil
}

func (m *MockRepository) GetQuestions(ctx context.Context, userId int, pages pagination.Pagination) (model.QuestionList, error) {
	list := make(model.QuestionList, 0)

	v, ok := m.order[userId]
	if !ok {
		return list, nil
	}

	start := pages.PageSize * pages.Page
	end := start + pages.PageSize

	if end >= len(v) {
		end = len(v) - 1
	}

	ids := v[start:end]
	for _, id := range ids {
		q, err := m.getQuestion(id, userId)
		if err != nil {
			return nil, err
		}

		list = append(list, *q)
	}

	return list, nil
}

func (m *MockRepository) UpdateQuestion(ctx context.Context, userId int, id int, updateFn func(q *model.Question) (*model.Question, error)) (*model.Question, error) {
	return nil, nil
}
