package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ResetPasswordInput struct {
	Email string `json:"email"`
}

func (input ResetPasswordInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.Email),
	)
}
