package store

import (
	"database/sql"
)

// Real example https://github.com/g0dm0d/blog-api/blob/main/store/user.go

type Example struct {
	ExampleRow sql.NullString
}

type ExampleFuncOpts struct {
	ExampleRow string
}

type ExampleStore interface {
	ExampleFunc(opts ExampleFuncOpts) (Example, error)
}
