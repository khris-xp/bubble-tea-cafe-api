package models

import "time"

type Cart struct {
	Id        string    `bson:"_id,omitempty"`
	UserId    string    `json:"user_id" bson:"user_id"`
	ProductId string    `json:"product_id" bson:"product_id"`
	Quantity  int       `json:"quantity" bson:"quantity"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
