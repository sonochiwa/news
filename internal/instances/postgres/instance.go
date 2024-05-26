package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sonochiwa/news/configs"
)

type Instance interface {
	Connect(postgres configs.Postgres) error
	Close() error
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	DB() *sql.DB
}

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}

func (p *Postgres) Connect(config configs.Postgres) error {
	connectionString := fmt.Sprintf(
		"host=%s port= %s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *Postgres) Close() error {
	if p.db == nil {
		return nil
	}

	return p.db.Close()
}

func (p *Postgres) Ping() error {
	if p.db == nil {
		return nil
	}

	return p.db.Ping()
}

func (p *Postgres) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if p.db == nil {
		return nil, errors.New("database: no connection")
	}

	return p.db.Query(query, args...)
}

func (p *Postgres) QueryRow(query string, args ...interface{}) *sql.Row {
	if p.db == nil {
		return nil
	}

	return p.db.QueryRow(query, args...)
}
