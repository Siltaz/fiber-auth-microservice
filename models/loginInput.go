package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input LoginInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.Email),
		validation.Field(&input.Password, validation.Required, validation.Length(8, 18)),
	)
}
