package models

import (
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	ProductName  string             `bson:"productName"`
	ProductPrice float32            `bson:"productPrice"`
}

func (model *ProductModel) ToProductDTO() *dto.ProductDTO {
	return &dto.ProductDTO{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		ProductName:  model.ProductName,
		ProductPrice: model.ProductPrice,
	}
}
