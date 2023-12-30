package repositories

import (
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	ProductCollection *mongo.Collection
}

func (repository *ProductRepository) GetAllProduct()  {}
func (repository *ProductRepository) GetProductById() {}

func (repository *ProductRepository) CreateProduct(payload *models.ProductModel) (*models.ProductModel, error) {
	result, err := repository.ProductCollection.InsertOne(ctx, payload)
	payload.ID = result.InsertedID.(primitive.ObjectID)
	return payload, err
}

func (repository *ProductRepository) UpdateProduct() {}
