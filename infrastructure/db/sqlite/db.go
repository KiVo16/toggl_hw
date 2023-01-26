package db

import (
	"base/internal/domain/model"
	"base/pkg/pagination"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type SQLiteDB struct {
	file string

	db *bun.DB
}

func NewSQLiteDB(ctx context.Context, opts ...SQLiteDBOption) (*SQLiteDB, error) {

	const (
		defaultFile = "test.db"
	)

	db := &SQLiteDB{
		file: defaultFile,
	}

	for _, opt := range opts {
		opt(db)
	}

	err := db.connect()
	if err != nil {
		return nil, err
	}

	go db.contextLoop(ctx)

	return db, nil
}

func (db *SQLiteDB) contextLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			err := db.close()
			if err != nil {
				log.Printf("db close err: %v\f", err)
				return
			}
			log.Println("db closed")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("DB not closed")
		}
	}
}

func (db *SQLiteDB) connect() error {
	sqldb, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	log.Println("test")

	db.db = bun.NewDB(sqldb, sqlitedialect.New())
	return nil
}

func (db *SQLiteDB) close() error {
	return db.db.Close()
}

func (db SQLiteDB) CreateQuestion(ctx context.Context, question model.Question) error {
	q := NewQuestion(question)

	err := db.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

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
	})
	return err
}

func (db SQLiteDB) DeleteQuestionByID(ctx context.Context, id int) error {
	q := Question{ID: int64(id)}

	_, err := db.db.NewDelete().
		Model(&q).
		WherePK().
		Exec(ctx)

	return err
}

func (db SQLiteDB) GetQuestions(ctx context.Context, pages pagination.Pagination) (model.QuestionList, error) {
	list := make([]Question, 0)

	err := db.db.NewSelect().
		Model(&list).
		Relation("Options").
		Limit(pages.Limit).
		Offset(pages.Offset).
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

func (db SQLiteDB) UpdateQuestion(ctx context.Context, id int, updateFn func(q *model.Question) (*model.Question, error)) error {
	return nil
}
