package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required,min=2,max=100" bson:"name"`
	Category   string             `json:"category" validate:"required,min=2,max=100" bson:"category"`
	Start_Date *time.Time         `json:"start_date" validate:"required" bson:"start_date"`
	End_Date   *time.Time         `json:"end_date" validate:"required" bson:"end_date"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
	Updated_At time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Menu_Id    string             `json:"menu_id" validate:"required" bson:"menu_id"`
}

