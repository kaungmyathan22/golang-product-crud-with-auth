package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/controllers"
)

func InitAuthenticationRoutes(routeGroup fiber.Router) {
	authentication_controller := controllers.AuthenticationController{}
	router := routeGroup.Group("/authentication")
	router.Post("/login", authentication_controller.Login)
	router.Post("/logout", authentication_controller.Logout)
	router.Post("/forgot-password", authentication_controller.ForgotPassword)
	router.Post("/reset-password", authentication_controller.ResetPassword)
	router.Post("/change-password", authentication_controller.ChangePassword)
	router.Post("/verify-email", authentication_controller.VerifyEmail)
	router.Post("/refresh-token", authentication_controller.RefreshToken)
}
