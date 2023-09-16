package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID primitive.ObjectID `bson:"_id"`
	Quantity *string `json:"quantity" validade:"required,eq=S|eq=M|eq=L"`
	Unit_price *float64 `json:"unit_price" validade:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Food_id *string `json:"food_id" validade:"required"`
	Order_item_id string `json:"order_item_id"`
	Order_id string `json:"order_id" validade:"required"`
}