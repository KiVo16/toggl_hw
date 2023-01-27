package db

import (
	"context"

	"github.com/uptrace/bun"
)

func (db SQLiteDB) getQuestion(ctx context.Context, userId, id int64) (*Question, error) {
	q := new(Question)

	err := db.db.NewSelect().
		Model(q).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Relation("Options").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return q, nil
}

func (db SQLiteDB) insertQuestionTX(ctx context.Context, tx bun.Tx, q *Question) error {
	_, err := tx.NewInsert().
		Model(q).
		Exec(ctx)

	if err != nil {
		return err
	}

	if q.Options != nil {
		q.SyncOptionsQuestionIDs()

		_, err := tx.NewInsert().
			Model(&q.Options).
			Exec(ctx)
		return err
	}
	return nil
}
