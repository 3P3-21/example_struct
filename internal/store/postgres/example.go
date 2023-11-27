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

// Real example https://github.com/g0dm0d/blog-api/blob/main/store/postgres/user.go#L26-L34
func (s *ExampleStore) ExampleFunc(opts store.ExampleFuncOpts) (store.Example, error) {
	return store.Example{}, nil
}
