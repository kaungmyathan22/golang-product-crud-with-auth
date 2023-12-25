package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/config"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
)

func main() {
	config.BootstrapApp()
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "pong",
		})
	})

	host := "localhost"
	port := fmt.Sprintf(":%s", config.AppConfigInstance.PORT)
	logger.Info(fmt.Sprintf("Server is running at http://%s%s/", host, port))
	err := app.Listen(port)
	logger.Fatal(err.Error())
}
