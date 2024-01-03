package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	TokenHash      string             `bson:"tokenHash"`
	ExpirationTime time.Time          `bson:"expirationTime"`
}
