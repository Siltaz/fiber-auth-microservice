package services

import (
	"context"
	"dmb-auth-service/global"
	"dmb-auth-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (models.User, error) {
	res, err := global.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
		return models.User{}, err
	}
	user, _ = GetUserByID(res.InsertedID.(primitive.ObjectID))
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

func GetUserByID(uid primitive.ObjectID) (models.User, error) {
	var user models.User
	err := global.DB.Collection("users").FindOne(context.Background(), bson.M{"_id": uid}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
