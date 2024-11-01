package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

var dbconn *sql.DB

func Init() error {
	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	dbFilePath := filepath.Join(rootDir, "datastore", "todos.db")

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	dbconn = db

	return nil
}

func GetConnection() *sql.DB {
	return dbconn
}

func Close() error {
	return dbconn.Close()
}
