package database

import (
	"database/sql"
	"errors"
)

var (
	ErrNotConnected = errors.New("database: no connection")
)

type Database interface {
	Connect() error
	Close() error
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
