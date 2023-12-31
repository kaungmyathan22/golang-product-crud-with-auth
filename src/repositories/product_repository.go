package repositories

import (
	"fmt"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/common"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	ProductCollection *mongo.Collection
}

func (repository *ProductRepository) GetAllProductCount(filter bson.M) (int64, error) {
	count, err := repository.ProductCollection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repository *ProductRepository) GetAllProduct(params *common.PaginationParams) ([]*dto.ProductDTO, error) {
	skip := (params.Page - 1) * params.PageSize

	cursor, err := repository.ProductCollection.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: &params.PageSize,
		Skip:  &skip,
	})
	if err != nil {
		return nil, err
	}
	var products []*dto.ProductDTO
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

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

func (repository *ProductRepository) UpdateProduct(payload *dto.UpdateProductDTO, filter *bson.M) (*dto.ProductDTO, error) {
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := repository.ProductCollection.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: payload}}, options)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var updatedProduct dto.ProductDTO
	if err := result.Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

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
