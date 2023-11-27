package dto

import "github.com/3P3-21/curriculum/model"

type Example struct {
	ExapleRow string `json:"example_row"`
}

func NewExample(m model.Example) Example {
	return Example{
		ExapleRow: m.ExapleRow,
	}
}

func NewExamples(m []model.Example) []Example {
	var examples []Example
	for i := range m {
		examples = append(examples, NewExample(m[i]))
	}
	return examples
}
