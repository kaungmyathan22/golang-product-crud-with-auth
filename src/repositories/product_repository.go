package repositories

import (
	"fmt"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	ProductCollection *mongo.Collection
}

func (repository *ProductRepository) GetAllProduct() {}

func (repository *ProductRepository) GetProductById(productId primitive.ObjectID) (*dto.ProductDTO, error) {
	var product dto.ProductDTO
	err := repository.ProductCollection.FindOne(ctx, bson.M{"_id": productId}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("product not found with given id %s", productId.Hex())
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *ProductRepository) CreateProduct(payload *models.ProductModel) (*models.ProductModel, error) {
	result, err := repository.ProductCollection.InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}
	payload.ID = result.InsertedID.(primitive.ObjectID)
	return payload, err
}

func (repository *ProductRepository) UpdateProduct() {}

func (repository *ProductRepository) DeleteProduct(productId primitive.ObjectID) error {

	result, err := repository.ProductCollection.DeleteOne(ctx, bson.M{"_id": productId})
	if result.DeletedCount == 0 {
		return fmt.Errorf("product not found with given id %s", productId.Hex())
	}
	if err != nil {
		return err
	}
	return nil
}
