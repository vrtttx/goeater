package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID 	primitive.ObjectID `bson:"_id"`
	Name *string `json:"name" validade:"required,min=2,max=100"`
	Price *float64 `json:"price" validade:"required"`
	Food_image *string `json:"food_image" validade:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Food_id string `json:"food_id"`
	Menu_id *string `json:"menu_id" validade:"required"`
}