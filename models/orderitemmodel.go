package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L" bson:"quantity"`
	Unit_Price    *float64           `json:"unit_price" validate:"required,min=0" bson:"unit_price"`
	Created_At    time.Time          `json:"created_at" bson:"created_at"`
	Updated_At    time.Time          `json:"updated_at" validate:"required" bson:"updated_at"`
	Food_Id       *string            `json:"food_id" validate:"required" bson:"food_id"`
	Order_Item_Id string             `json:"order_item_id" validate:"required" bson:"order_item_id"`
	Order_Id      string             `json:"order_id" validate:"required" bson:"order_id"`
}
