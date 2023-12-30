package repositories

import (
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

func (repository *ProductRepository) GetProductById(productId string) (*dto.ProductDTO, error) {
	var product dto.ProductDTO
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, err
	}
	err = repository.ProductCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&product)
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
