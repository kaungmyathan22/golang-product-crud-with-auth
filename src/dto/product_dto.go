package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDTO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ProductName  string             `bson:"productName" json:"productName"`
	ProductPrice float32            `bson:"productPrice" json:"productPrice"`
}

type CreateProductDTO struct {
	ProductName  string  `bson:"productName" json:"productName"`
	ProductPrice float32 `bson:"productPrice" json:"productPrice"`
}

type UpdateProductDTO struct {
	ProductName  string  `bson:"productName" json:"productName"`
	ProductPrice float32 `bson:"productPrice" json:"productPrice"`
}
