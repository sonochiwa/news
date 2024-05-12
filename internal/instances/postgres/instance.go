package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"news/internal/configs"
)

type Instance interface {
	Connect(postgres configs.Postgres) error
	Close() error
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type pgInstance struct {
	db *sql.DB
}

func (p *pgInstance) Connect(config configs.Postgres) error {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Username, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *pgInstance) Close() error {
	if p.db == nil {
		return nil
	}

	return p.db.Close()
}

func (p *pgInstance) Ping() error {
	if p.db == nil {
		return nil
	}

	return p.db.Ping()
}

func (p *pgInstance) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if p.db == nil {
		return nil, errors.New("database: no connection")
	}

	return p.db.Query(query, args...)
}

func (p *pgInstance) QueryRow(query string, args ...interface{}) *sql.Row {
	if p.db == nil {
		return nil
	}

	return p.db.QueryRow(query, args...)
}
