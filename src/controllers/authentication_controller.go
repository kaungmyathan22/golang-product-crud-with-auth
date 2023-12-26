package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type AuthenticationController struct {
	Service *services.UserService
}

func (controller *AuthenticationController) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Login route"})
}

func (controller *AuthenticationController) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Logout route"})
}

func (controller *AuthenticationController) Register(ctx *fiber.Ctx) error {
	var payload dto.CreateUserDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid payload",
			"error":   err.Error(),
		})
	}
	if err := validate.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}

	return ctx.JSON(fiber.Map{"message": "Register route"})
}

func (controller *AuthenticationController) RefreshToken(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "RefreshToken route"})
}

func (controller *AuthenticationController) ForgotPassword(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "ForgotPassword route"})
}

func (controller *AuthenticationController) ResetPassword(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "ResetPassword route"})
}

func (controller *AuthenticationController) VerifyEmail(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "VerifyEmail route"})
}

func (controller *AuthenticationController) ChangePassword(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "ChangePassword route"})
}
