package dto

type CreateRefreshTokenDTO struct {
	TokenHash      string `bson:"tokenHash" json:"tokenHash"`
	ExpirationTime string `bson:"expirationTime" json:"expirationTime"`
	UserID         string `bson:"userID" json:"userID"`
}
