package repository

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	DbClient *sql.DB
)

func DbConnector() error {
	connectionString := "postgresql://postgres:postgres@localhost:5432/school?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	DbClient = db
	return nil
}