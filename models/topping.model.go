package models

import (
	"time"
)

type Topping struct {
	Id        string    `bson:"_id,omitempty"`
	Name      string    `json:"name" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
