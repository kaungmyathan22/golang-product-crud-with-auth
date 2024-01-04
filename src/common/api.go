package common

import (
	"github.com/gofiber/fiber/v2"
)

func ParsePaginationParams(c *fiber.Ctx) *PaginationParams {
	params := new(PaginationParams)
	params.Page = 1
	params.PageSize = 10
	if err := c.QueryParser(params); err != nil {
		return params
	}

	// Set default values if the parameters are not present
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}

	return params
}

func InvalidPayloadErrorResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "invalid payload",
		"error":   err.Error(),
	})
}

func BadRequestErrorResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Code:   fiber.StatusBadRequest,
		Errors: TransformError(err.Error()),
	})
}
