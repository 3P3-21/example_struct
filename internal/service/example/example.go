package example

import (
	"github.com/3P3-21/curriculum/dto"
	"github.com/3P3-21/curriculum/internal/server/req"
	"github.com/3P3-21/curriculum/internal/store"
	"github.com/3P3-21/curriculum/model"
	"github.com/3P3-21/curriculum/pkg/errs"
)

func (s *Service) Example(ctx *req.Ctx) error {
	data, err := s.exampleStore.ExampleFunc(store.ExampleFuncOpts{
		ExampleRow: "test",
	})
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}
	return ctx.JSON(dto.NewExample(model.NewExample(data)))
}
