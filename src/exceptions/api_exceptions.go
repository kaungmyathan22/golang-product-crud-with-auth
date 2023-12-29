package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
)

func UnauthorizedRequestException() *common.ErrorResponse {
	return &common.ErrorResponse{
		Code: fiber.StatusUnauthorized,
		Errors: []string{
			"Unauthorized access need to login first.",
		},
	}
}
