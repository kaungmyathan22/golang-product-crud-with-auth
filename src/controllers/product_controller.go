package controllers

import "github.com/gofiber/fiber/v2"

type ProductController struct{}

func (controller *ProductController) GetProducts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Get all product endpoints"})
}

func (controller *ProductController) GetProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Get product endpoints"})
}

func (controller *ProductController) CreateProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create product endpoints"})
}

func (controller *ProductController) UpdateProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Update product endpoints"})
}

func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Delete product endpoints"})
}
