package model

import "github.com/3P3-21/curriculum/internal/store"

type Example struct {
	ExapleRow string
}

func NewExample(s store.Example) Example {
	example := Example{
		ExapleRow: s.ExampleRow.String,
	}

	return example
}

func NewExamples(s []store.Example) []Example {
	var examples []Example
	for i := range s {
		examples = append(examples, NewExample(s[i]))
	}
	return examples
}
