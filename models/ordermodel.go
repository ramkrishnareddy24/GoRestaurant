package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date *time.Time         `json:"order_date" validate:"required" bson:"order_date"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
	Updated_At time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Order_Id   string             `json:"order_id" validate:"required" bson:"order_id"`
	Table_Id   string             `json:"table_id" validate:"required" bson:"table_id"`
}
