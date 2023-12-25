package controllers

import "github.com/gofiber/fiber/v2"

type AuthenticationController struct {
}

func (controller *AuthenticationController) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Login route"})
}

func (controller *AuthenticationController) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Logout route"})
}

func (controller *AuthenticationController) Register(ctx *fiber.Ctx) error {
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
