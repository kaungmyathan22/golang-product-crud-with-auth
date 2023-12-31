package controllers

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/services"
)

type ProductController struct {
	ProductService *services.ProductService
}

func (controller *ProductController) GetProducts(ctx *fiber.Ctx) error {
	paginationParams := common.ParsePaginationParams(ctx)
	products, err := controller.ProductService.GetProductByProducts(paginationParams)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}
	count, err := controller.ProductService.GetProductsCount()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}
	totalPages := int64(math.Ceil(float64(count) / float64(paginationParams.PageSize)))
	var previousPage *int64
	if paginationParams.Page > 1 {
		temp := paginationParams.Page - 1
		previousPage = &temp
	} else {
		previousPage = nil
	}

	var nextPage *int64
	if paginationParams.Page+1 < totalPages {
		temp := paginationParams.Page + 1
		nextPage = &temp
	} else {
		nextPage = nil
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"page":         paginationParams.Page,
		"pageSize":     paginationParams.PageSize,
		"totalItems":   count,
		"totalPages":   totalPages,
		"nextPage":     nextPage,
		"previousPage": previousPage,
		"items":        products,
	})
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
	productId := ctx.Params("id")
	var payload *dto.UpdateProductDTO
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

	product, err := controller.ProductService.UpdateProduct(payload, productId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Errors: common.TransformError(err.Error()),
		})
	}

	return ctx.JSON(fiber.Map{"data": product})
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
