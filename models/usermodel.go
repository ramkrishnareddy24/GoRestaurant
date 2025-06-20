package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_Name    *string            `json:"last_name" validate:"required,min=2,max=100" bson:"last_name"`
	Password      *string            `json:"password" validate:"required,min=6,max=100" bson:"password"`
	Email         *string            `json:"email" validate:"required,email" bson:"email"`
	Avatar        *string            `json:"avatar" validate:"omitempty,url" bson:"avatar"`
	Phone         *string            `json:"phone" validate:"required,min=10,max=15" bson:"phone"`
	Token         *string            `json:"token" bson:"token"`
	Refresh_Token *string            `json:"refresh_token" bson:"refresh_token"`
	Created_At    time.Time          `json:"created_at" bson:"created_at"`
	Updated_At    time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	User_Id       string             `json:"user_id" validate:"required" bson:"user_id"`
}
