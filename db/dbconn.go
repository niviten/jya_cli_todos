package db

import "database/sql"

func GetDB() (*sql.DB, error) {
    return sql.Open("sqlite3", "db/database.db")
}
