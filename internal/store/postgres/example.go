package postgres

import (
	"database/sql"

	"github.com/3P3-21/curriculum/internal/store"
)

type ExampleStore struct {
	db *sql.DB
}

func NewExampleStore(db *sql.DB) store.ExampleStore {
	return &ExampleStore{
		db: db,
	}
}

func (s *ExampleStore) ExampleFunc(opts store.ExampleFuncOpts) (store.Example, error) {
	return store.Example{}, nil
}
