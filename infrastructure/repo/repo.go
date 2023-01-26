package repo

import (
	db "base/infrastructure/db/sqlite"
	"base/internal/interfaces"
)

func NewRepoWithSQLite(db db.SQLiteDB) interfaces.Repository {
	return db
}
