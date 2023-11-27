package store

import (
	"database/sql"
)

type Example struct {
	ExampleRow sql.NullString
}

type ExampleFuncOpts struct {
	ExampleRow string
}

type ExampleStore interface {
	ExampleFunc(opts ExampleFuncOpts) (Example, error)
}
