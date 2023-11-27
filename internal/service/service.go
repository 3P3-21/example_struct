package service

import (
	"github.com/3P3-21/curriculum/internal/service/example"
	"github.com/3P3-21/curriculum/internal/store"
)

type Service struct {
	Example example.Example
}

func New(exampleStore store.ExampleStore) *Service {
	return &Service{
		Example: example.New(exampleStore),
	}
}
