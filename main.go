package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/middlewares"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/routes"
)

func main() {
	config.BootstrapApp()
	client := config.ConnectToDatabase()
	app := fiber.New()
	v1Group := app.Group("/api/v1", middlewares.LoggerMiddleware)
	routes.InitAuthenticationRoutes(v1Group, client)
	routes.InitProductRoutes(v1Group, client)
	v1Group.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "pong",
		})
	})

	//#region------------------- app config & run section
	host := "localhost"
	port := fmt.Sprintf(":%s", config.AppConfigInstance.PORT)
	logger.Info(fmt.Sprintf("Server is running at http://%s%s/", host, port))
	err := app.Listen(port)
	logger.Fatal(err.Error())
	//#endregion

}
