package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordResetTokenModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UserID         primitive.ObjectID `bson:"userID"`
	Code           string             `bson:"code"`
	ExpirationTime time.Time          `bson:"expirationTime"`
}
