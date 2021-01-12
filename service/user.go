package service

import (
	"context"
	"dmb-auth-service/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"-" bson:"updatedAt"`
	DeletedAt *time.Time         `json:"-" bson:"deletedAt"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 18)),
	)
}

func CreateUser(user User) (User, error) {
	res, err := config.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
		return User{}, err
	}
	user, _ = GetUserByID(res.InsertedID.(primitive.ObjectID))
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(id primitive.ObjectID) (User, error) {
	var user User
	err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
