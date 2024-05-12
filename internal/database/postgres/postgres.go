package postgres

import (
	"database/sql"
	"news/internal/database"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=news sslmode=disable")
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
		return nil, database.ErrNotConnected
	}

	return p.db.Query(query, args...)
}

func (p *Postgres) QueryRow(query string, args ...interface{}) *sql.Row {
	if p.db == nil {
		return nil
	}

	return p.db.QueryRow(query, args...)
}

func New() (database.Database, error) {
	psql := &Postgres{}

	err := psql.Connect()
	if err != nil {
		return nil, err
	}

	err = psql.Ping()
	if err != nil {
		psql.Close()
		return nil, err
	}

	return psql, nil
}
