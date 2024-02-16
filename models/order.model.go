package models

import (
	"time"
)

type Order struct {
	Id        string    `bson:"_id,omitempty"`
	UserId    string    `json:"user_id" bson:"user_id" validate:"required"`
	MenuId    string    `json:"menu_id" bson:"menu_id" validate:"required"`
	Topping   []string  `json:"topping" bson:"topping" validate:"required"`
	Quantity  int       `json:"quantity" bson:"quantity" validate:"required"`
	Total     float64   `json:"total" bson:"total" validate:"required"`
	Status    string    `json:"status" bson:"status" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
