package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/controllers"
)

func InitProductRoutes(routeGroup fiber.Router) {
	product_controller := controllers.ProductController{}
	router := routeGroup.Group("/products")
	router.Get("/", product_controller.GetProducts)
	router.Post("/", product_controller.CreateProduct)
	router.Get("/:id", product_controller.GetProduct)
	router.Patch("/:id", product_controller.UpdateProduct)
	router.Delete("/:id", product_controller.DeleteProduct)
}
