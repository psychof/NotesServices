package pkg

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func New() *Validator {
	return &Validator{Validator: validator.New()}
}

func (c *Validator) Validate(s interface{}) error {
	return c.Validator.Struct(s)
}
