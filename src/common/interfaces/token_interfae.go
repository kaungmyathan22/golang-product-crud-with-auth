package interfaces

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	jwt.StandardClaims
}
