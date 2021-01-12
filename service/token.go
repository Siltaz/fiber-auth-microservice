package service

import (
	"context"
	"dmb-auth-service/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthUUID  string             `bson:"auth_uuid"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	ExpireAt  time.Time          `bson:"expireAt"`
}

func CreateRefreshToken(token RefreshToken) (RefreshToken, error) {
	res, err := config.DB.Collection("refresh_tokens").InsertOne(context.Background(), token)
	if err != nil {
		panic(err)
		return RefreshToken{}, err
	}
	token, _ = GetRefreshTokenByID(res.InsertedID.(primitive.ObjectID))
	return token, nil
}

func GetRefreshTokenByID(id primitive.ObjectID) (RefreshToken, error) {
	var token RefreshToken
	err := config.DB.Collection("refresh_tokens").FindOne(context.Background(), bson.M{"_id": id}).Decode(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}
