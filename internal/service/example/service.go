package example

import (
	"github.com/3P3-21/curriculum/internal/server/req"
	"github.com/3P3-21/curriculum/internal/store"
)

type Example interface {
	Example(ctx *req.Ctx) error
}

type Service struct {
	exampleStore store.ExampleStore
}

func New(exampleStore store.ExampleStore) *Service {
	return &Service{
		exampleStore: exampleStore,
	}
}
