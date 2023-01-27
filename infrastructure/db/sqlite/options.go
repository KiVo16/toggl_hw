package db

type SQLiteDBOption func(*SQLiteDB)

func WithFile(file string) SQLiteDBOption {
	return func(db *SQLiteDB) {
		db.file = file
	}
}

func WithMigrationsPath(migrationsPath string) SQLiteDBOption {
	return func(db *SQLiteDB) {
		db.migrationsPath = migrationsPath
	}
}
