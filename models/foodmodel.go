package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100" bson:"name"`
	Price      float64            `json:"price" validate:"required" bson:"price"`
	FoodImage  *string            `json:"food_image" validate:"required" bson:"food_image"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
	Updated_At time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Food_Id    *string            `json:"food_id" validate:"required" bson:"food_id"`
	Menu_Id    *string            `json:"menu_id" validate:"required" bson:"menu_id"`
}

