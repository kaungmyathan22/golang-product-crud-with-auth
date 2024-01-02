package routes

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/controllers"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/middlewares"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitAuthenticationRoutes(routeGroup fiber.Router, client *mongo.Client) {
	userCollection := client.Database(config.AppConfigInstance.DATABASE_NAME).Collection(config.AppConfigInstance.USER_COLLECTION)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := userCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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

	router.Post("/forgot-password", authentication_controller.ForgotPassword)
	router.Post("/reset-password", authentication_controller.ResetPassword)
	router.Post("/verify-email", authentication_controller.VerifyEmail)
	router.Post("/refresh-token", authentication_controller.RefreshToken)

	router.Get("/logout", middlewares.IsAuthenticatedMiddleware(authentication_service), authentication_controller.Logout)
	router.Get("/me", middlewares.IsAuthenticatedMiddleware(authentication_service), authentication_controller.Me)
	router.Get("/delete-account", middlewares.IsAuthenticatedMiddleware(authentication_service), authentication_controller.DeleteAccount)
	router.Post("/change-password", middlewares.IsAuthenticatedMiddleware(authentication_service), authentication_controller.ChangePassword)

}
