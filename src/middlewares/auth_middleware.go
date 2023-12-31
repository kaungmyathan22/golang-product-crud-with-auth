package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common/interfaces"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsAuthenticatedMiddleware(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationValue := c.Get("Authorization")
		splittedAuthHeader := strings.Split(authorizationValue, "Bearer ")
		tokenString := ""
		if len(splittedAuthHeader) > 1 {
			tokenString = splittedAuthHeader[1]
		}
		if authorizationValue == "" || tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Errors: []string{"Authorization token is required."},
			})
		}
		claims := &interfaces.JwtCustomClaims{}
		secretKey := []byte(config.AppConfigInstance.ACCESS_TOKEN_SECRET)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
				Code:   fiber.StatusBadRequest,
				Errors: []string{"Invalid authorization token."},
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Errors: []string{"Invalid authorization token."},
			})
		}

		expires := claims.Exp
		if time.Now().Unix() > expires {
			return c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Errors: []string{"Authorization token expired."},
			})
		}
		userID := claims.Sub
		user, err := userService.GetUserByUserId(userID)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
					Code:   fiber.StatusUnauthorized,
					Errors: common.TransformError("User not found associating with this token. May be account has been deleted."),
				})
			}
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}
