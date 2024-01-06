package dto

import (
	"time"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=3" error:"username field is required and must be a minimum of 3."`
	Password string `json:"password" validate:"required,min=6" error:"password field is required and must be a minimum of 6."`
}

type CreatePasswordResetDTO struct {
	UserID         string    `bson:"userID"`
	Code           string    `bson:"code"`
	ExpirationTime time.Time `bson:"expirationTime"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" validate:"required,email" error:"email field is required and must be a valid email address."`
}

type PasswordResetCodeConfirmationDTO struct {
	Code string `json:"code" validate:"required,min=6,max=6"`
}
