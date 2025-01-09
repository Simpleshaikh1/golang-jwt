package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     *string            `json:"first_name" validate:"required, min=2, max=100"`
	LastName      *string            `json:"last_name" validate:"required, min=2, max=100"`
	Password      *string            `json:"password" validate:"required,min=8"`
	Email         *string            `json:"email" validate:"required,email"`
	Phone         *string            `json:"phone" validate:"required,number"`
	Token         *string            `json:"token" validate:"required,min=8"`
	User_type     *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       *string
}
