package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"golang.org/x/crypto/scrypt"
)

type TokenService struct {
	Repository *repositories.TokenRepository
}

func signToken(expTime time.Time, userId string, secretKey []byte) (string, error) {
	claims := interfaces.JwtCustomClaims{
		Sub: userId,
		Iat: time.Now().Unix(),
		Exp: expTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tokenService *TokenService) SignAccessToken(userId string) (string, error) {
	return signToken(tokenService.GetAccessTokenExpiration(), userId, []byte(config.AppConfigInstance.ACCESS_TOKEN_SECRET))
}

func (tokenService *TokenService) GetAccessTokenExpiration() time.Time {
	return time.Now().Add(1 * time.Hour)
}

func (tokenService *TokenService) GetRefreshTokenExpiration() time.Time {
	return time.Now().Add(7 * 24 * time.Hour)
}

func (tokenService *TokenService) SignRefreshToken(userId string) (string, error) {

	token, err := signToken(tokenService.GetRefreshTokenExpiration(), userId, []byte(config.AppConfigInstance.REFRESH_TOKEN_SECRET))
	if err != nil {
		return "", err
	}
	// hash token before saving to database
	hashedToken, err := tokenService.EncryptRefreshToken(token)
	if err != nil {
		return "", err
	}
	payload := dto.CreateRefreshTokenDTO{
		UserID:         userId,
		ExpirationTime: tokenService.GetRefreshTokenExpiration(),
		TokenHash:      string(hashedToken),
	}
	if err = tokenService.Repository.CreateNewToken(&payload); err != nil {
		return "", err
	}
	return token, nil
}

func (tokenService *TokenService) EncryptRefreshToken(token string) (string, error) {
	key := make([]byte, 32) // Use 32 bytes as an example, adjust as needed

	copy(key, []byte(config.AppConfigInstance.REFRESH_TOKEN_ENCRYPT_KEY))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(token), nil)
	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}

func DecryptRefreshToken(encryptedToken string) (string, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(encryptedToken)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(config.AppConfigInstance.REFRESH_TOKEN_ENCRYPT_KEY))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext is too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	decryptedToken, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedToken), nil
}

func GenerateAESKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 32768, 8, 1, 32)
}
