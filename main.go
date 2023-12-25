package main

import (
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
	logger.Info("Server is running at http://localhost:8000/")
	err := app.Listen(":8000")
	logger.Fatal(err.Error())
}
