package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_Of_Guests *int               `json:"number_of_guests" validate:"required,min=1,max=20" bson:"number_of_guests"`
	Table_Number     *int               `json:"table_number" validate:"required,min=1,max=100" bson:"table_number"`
	Created_At       time.Time          `json:"created_at" bson:"created_at"`
	Updated_At       time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Table_Id         string             `json:"table_id" validate:"required" bson:"table_id"`
}
