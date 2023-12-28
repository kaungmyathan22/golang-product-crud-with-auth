package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/controllers"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAuthenticationRoutes(routeGroup fiber.Router, client *mongo.Client) {
	userCollection := client.Database(config.AppConfigInstance.DATABASE_NAME).Collection(config.AppConfigInstance.USER_COLLECTION)
	userRepository := &repositories.UserRepository{
		UserCollection: userCollection,
	}
	authentication_service := services.UserService{
		Repository: userRepository,
	}
	authentication_controller := controllers.AuthenticationController{
		Service:      &authentication_service,
		TokenService: &services.TokenService{},
	}
	router := routeGroup.Group("/authentication")
	router.Post("/login", authentication_controller.Login)
	router.Post("/register", authentication_controller.Register)
	router.Post("/logout", authentication_controller.Logout)
	router.Post("/forgot-password", authentication_controller.ForgotPassword)
	router.Post("/reset-password", authentication_controller.ResetPassword)
	router.Post("/change-password", authentication_controller.ChangePassword)
	router.Post("/verify-email", authentication_controller.VerifyEmail)
	router.Post("/refresh-token", authentication_controller.RefreshToken)
}
