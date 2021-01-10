package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name" validate:"required,min=3,max=64"`
	Email    string             `json:"email" validate:"required,email,min=6,max=64"`
	Password string             `json:"password" validate:"required,min=8,max=18"`
}

type LoginInput struct {
	Email string `json:"email" validate:"required,email,min=6,max=64"`
	Password string `json:"password" validate:"required,min=8,max=18"`
}

type ResetPasswordInput struct {
	Email string `json:"email" validate:"required,email,min=6,max=64"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
