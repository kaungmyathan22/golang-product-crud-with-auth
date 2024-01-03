package dto

import "time"

type CreateRefreshTokenDTO struct {
	TokenHash      string    `bson:"tokenHash" json:"tokenHash"`
	ExpirationTime time.Time `bson:"expirationTime" json:"expirationTime"`
	UserID         string    `bson:"userID" json:"userID"`
}
