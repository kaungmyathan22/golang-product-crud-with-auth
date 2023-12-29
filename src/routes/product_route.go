package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/controllers"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/middlewares"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitProductRoutes(routeGroup fiber.Router, client *mongo.Client) {
	product_controller := controllers.ProductController{}
	userCollection := client.Database(config.AppConfigInstance.DATABASE_NAME).Collection(config.AppConfigInstance.USER_COLLECTION)
	userRepository := &repositories.UserRepository{
		UserCollection: userCollection,
	}
	authentication_service := services.UserService{
		Repository: userRepository,
	}
	router := routeGroup.Group("/products", middlewares.IsAuthenticatedMiddleware(authentication_service))
	router.Get("/", product_controller.GetProducts)
	router.Post("/", product_controller.CreateProduct)
	router.Get("/:id", product_controller.GetProduct)
	router.Patch("/:id", product_controller.UpdateProduct)
	router.Delete("/:id", product_controller.DeleteProduct)
}
