package services

import (
	"context"
	"dmb-auth-service/global"
	"dmb-auth-service/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user models.User) (models.User, error) {
	res, err := global.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
		return models.User{}, err
	}
	fmt.Println(res.InsertedID)
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := global.DB.Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
