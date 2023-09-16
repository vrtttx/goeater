package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	First_name *string `json:"first_name" validade:"required,min=2,max=100"`
	Last_name *string `json:"last_name" validade:"required,min=2,max=100"`
	Password *string `json:"password" validate:"required,min=8"`
	Email *string `json:"email" validade:"email,required"`
	Avatar *string `json:"avatar"`
	Phone *string `json:"phone" validade:"required"`
	Token *string `json:"token"`
	Refresh_Token *string `json:"refresh_token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id string `json:"user_id"`
}