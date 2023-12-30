package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
)

type ProductController struct {
	ProductService *services.ProductService
}

func (controller *ProductController) GetProducts(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Get all product endpoints"})
}

func (controller *ProductController) GetProduct(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")
	product, err := controller.ProductService.GetProductByProductId(productId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: []string{err.Error()},
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"data": product})
}

func (controller *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var payload *dto.CreateProductDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(common.ErrorResponse{
			Errors: common.TransformError(err.Error()),
			Code:   fiber.StatusUnprocessableEntity,
		})
	}
	if err := validate.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}

	product, err := controller.ProductService.CreateProduct(payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}

	return ctx.JSON(fiber.Map{"data": product})
}

func (controller *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Update product endpoints"})
}

func (controller *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")
	err := controller.ProductService.DeleteProductByProductId(productId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: []string{err.Error()},
		})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Successfully deleted product."})
}
