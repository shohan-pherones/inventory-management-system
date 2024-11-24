package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Quantity int                `json:"quantity" bson:"quantity"`
	Price    float64            `json:"price" bson:"price"`
}
