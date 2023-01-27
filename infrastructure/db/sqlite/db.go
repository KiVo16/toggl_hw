package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		return nil, fmt.Errorf("[db] connect err: %v", err)
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
				log.Printf("[db] close err: %v\f", err)
				return
			}
			log.Println("[db] closed")
			return
		case <-time.After(1 * time.Second):
		}
	}
}

func (db *SQLiteDB) connect() error {
	_, err := os.Stat(db.file)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		_, err := os.Create(db.file)
		if err != nil {
			return err
		}
	}

	datasource := fmt.Sprintf("%s?_foreign_keys=on", db.file)
	sqldb, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return err
	}

	db.db = bun.NewDB(sqldb, sqlitedialect.New())

	log.Println("[db] connected")

	err = db.runMigration()
	if err != nil {
		return err
	}

	log.Println("[db] migration up")

	return nil
}

func (db *SQLiteDB) close() error {
	return db.db.Close()
}

func (db SQLiteDB) runMigration() error {
	migrationFiles := "file://infrastructure/db/sqlite/migrations"
	driver, err := sqlite3.WithInstance(db.db.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationFiles, "sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
