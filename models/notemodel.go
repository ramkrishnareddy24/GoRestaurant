package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text" validate:"required,min=1,max=500" bson:"text"`
	Title      string             `json:"title" validate:"required,min=1,max=100" bson:"title"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
	Updated_At time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Note_Id    string             `json:"note_id" validate:"required" bson:"note_id"`
}

