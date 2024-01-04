package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordResetTokenModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	UserID         primitive.ObjectID `bson:"userID"`
	Code           string             `bson:"code"`
	ExpirationTime time.Time          `bson:"expirationTime"`
}
