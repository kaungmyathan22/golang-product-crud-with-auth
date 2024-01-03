package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/exceptions"
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
	if !user.IsEmailVerified {
		return ctx.Status(fiber.StatusForbidden).JSON(common.ErrorResponse{
			Errors: common.TransformError("Please verify your email first."),
			Code:   fiber.StatusForbidden,
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
	refreshToken, err := controller.TokenService.SignRefreshToken(user.ID.Hex())
	if err != nil {
		logger.Info(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Errors: common.TransformError("error signing refresh token."),
			Code:   fiber.StatusInternalServerError,
		})
	}
	cookie := &fiber.Cookie{
		Name:     config.AppConfigInstance.ACCESS_TOKEN_COOKIE_NAME,
		Value:    token,
		Expires:  controller.TokenService.GetAccessTokenExpiration(),
		HTTPOnly: true, // Make the cookie HTTP-only for added security
		SameSite: "Strict",
	}
	ctx.Cookie(cookie)

	cookie = &fiber.Cookie{
		Name:     config.AppConfigInstance.REFRESH_TOKEN_COOKIE_NAME,
		Value:    refreshToken,
		Expires:  controller.TokenService.GetRefreshTokenExpiration(),
		HTTPOnly: true, // Make the cookie HTTP-only for added security
		SameSite: "Strict",
	}
	ctx.Cookie(cookie)
	return ctx.JSON(fiber.Map{"user": user})
}

func (controller *AuthenticationController) Logout(ctx *fiber.Ctx) error {
	expiredCookie := fiber.Cookie{
		Name:     config.AppConfigInstance.ACCESS_TOKEN_COOKIE_NAME,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Set expiration time in the past
		HTTPOnly: true,
		SameSite: "Strict",
	}
	// Set the expired cookie in the response
	ctx.Cookie(&expiredCookie)
	return ctx.JSON(fiber.Map{"message": "Successfully logout."})
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

func (controller *AuthenticationController) ResendVerificationEmail(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "ResendVerificationEmail"})
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
	var payload dto.ChangePasswordDTO
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

	if payload.NewPassword == payload.OldPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(("New password cannot be the same with old password.")),
		})
	}
	user, ok := ctx.Context().UserValue("user").(*dto.UserDTO)
	if !ok {
		ctx.Status(fiber.StatusUnauthorized).JSON(exceptions.UnauthorizedRequestException())
	}
	if err := user.IsPasswordMach(payload.OldPassword); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Code:   fiber.StatusUnauthorized,
			Errors: common.TransformError(("Incorrect old password.")),
		})
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), bcrypt.DefaultCost) //payload.Password
	if err != nil {
		logger.Info(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Errors: common.TransformError(("Something went wrong.")),
		})
	}
	if err = controller.Service.ChangePassword(user, string(hashed_password)); err != nil {
		logger.Info(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Code:   fiber.StatusInternalServerError,
			Errors: common.TransformError(("Error changing password.")),
		})
	}
	return ctx.JSON(fiber.Map{"message": "Successfully changed password. Please log back in."})
}

func (controller *AuthenticationController) Me(ctx *fiber.Ctx) error {
	user, ok := ctx.Context().UserValue("user").(*dto.UserDTO)
	if !ok {
		ctx.Status(fiber.StatusUnauthorized).JSON(exceptions.UnauthorizedRequestException())
	}
	return ctx.JSON(user)
}

func (controller *AuthenticationController) DeleteAccount(ctx *fiber.Ctx) error {
	user, ok := ctx.Context().UserValue("user").(*dto.UserDTO)
	if !ok {
		ctx.Status(fiber.StatusUnauthorized).JSON(exceptions.UnauthorizedRequestException())
	}
	err := controller.Service.DeleteUserById(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Successfully delete account.",
	})
}
