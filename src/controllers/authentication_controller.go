package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type AuthenticationController struct {
	Service      *services.UserService
	TokenService *services.TokenService
}

func (controller *AuthenticationController) Login(ctx *fiber.Ctx) error {
	var payload dto.LoginDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Errors: common.TransformError(err.Error()),
			Code:   fiber.StatusBadRequest,
		})
	}
	if err := validate.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}
	user, err := controller.Service.GetUserByUsername(payload.Username)
	if err != nil {
		if err := ctx.BodyParser(&payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
				Errors: common.TransformError(err.Error()),
				Code:   fiber.StatusBadRequest,
			})
		}
	}
	if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Errors: common.TransformError("invalid username / password."),
			Code:   fiber.StatusUnauthorized,
		})
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Errors: common.TransformError("invalid username / password."),
			Code:   fiber.StatusUnauthorized,
		})
	}
	token, err := controller.TokenService.SignAccessToken(user.ID.Hex())
	if err != nil {
		logger.Info(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Errors: common.TransformError("error signing access token."),
			Code:   fiber.StatusInternalServerError,
		})
	}
	return ctx.JSON(fiber.Map{"user": user, "token": token})
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
	user, err := controller.Service.CreateUser(&payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}

	return ctx.JSON(fiber.Map{"message": "Successfully registered.", "data": map[string]string{
		"username": user.Username,
	}})
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
