package models

import (
	"time"
)

type Menu struct {
	Id          string    `bson:"_id,omitempty"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" bson:"description" validate:"required"`
	Price       float64   `json:"price" bson:"price" validate:"required"`
	Category    string    `json:"category" bson:"category" validate:"required"`
	Image       string    `json:"image" bson:"image" validate:"required"`
	Toppings    []string  `json:"toppings" bson:"toppings"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
