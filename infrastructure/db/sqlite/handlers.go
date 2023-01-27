package db

import (
	"base/internal/domain/model"
	e "base/internal/errors"
	"base/pkg/pagination"
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

func (db SQLiteDB) CreateQuestion(ctx context.Context, question model.Question) (*model.Question, error) {
	q := NewQuestion(question)

	err := db.sql.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		return db.insertQuestionTX(ctx, tx, q)
	})

	if err != nil {
		return nil, err
	}

	qp := q.ToModel()
	return &qp, nil
}

func (db SQLiteDB) DeleteQuestionByID(ctx context.Context, userId, id int) error {
	q, err := db.getQuestion(ctx, int64(userId), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return e.ErrQuestionNotFound
		}
		return err
	}

	_, err = db.sql.NewDelete().
		Model(q).
		WherePK().
		Exec(ctx)

	return err
}

func (db SQLiteDB) GetQuestions(ctx context.Context, userId int, pages pagination.Pagination) (model.QuestionList, error) {
	list := make([]Question, 0)

	limit, offset := getLimitOffsetFromModelPagination(pages)

	err := db.sql.NewSelect().
		Model(&list).
		Where("user_id = ?", userId).
		Relation("Options").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	fList := make(model.QuestionList, len(list))
	for i, q := range list {
		fList[i] = q.ToModel()
	}

	return fList, nil
}

func (db SQLiteDB) UpdateQuestion(ctx context.Context, userId, id int, updateFn func(q *model.Question) (*model.Question, error)) (*model.Question, error) {
	q, err := db.getQuestion(ctx, int64(userId), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.ErrQuestionNotFound
		}
		return nil, err
	}

	modelQuestion := q.ToModel()
	_, err = updateFn(&modelQuestion)
	if err != nil {
		return nil, err
	}

	q = NewQuestion(modelQuestion)

	err = db.sql.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		_, err := db.sql.NewDelete().
			Model(q).
			WherePK().
			Exec(ctx)

		if err != nil {
			return err
		}

		return db.insertQuestionTX(ctx, tx, q)
	})

	if err != nil {
		return nil, err
	}

	qp := q.ToModel()
	return &qp, err
}
