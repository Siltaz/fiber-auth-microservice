package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	Email     string              `json:"email" bson:"email"`
	Password  string              `json:"-" bson:"password"`
	CreatedAt primitive.DateTime  `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime  `json:"-" bson:"updated_at"`
	DeletedAt *primitive.DateTime `json:"-" bson:"deleted_at"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 18)),
	)
}
