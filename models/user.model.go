package models

import (
	"time"
)

type User struct {
	Id        string    `bson:"_id,omitempty"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Role      string    `json:"role" bson:"role"`
	Cart      []Cart    `json:"cart" bson:"cart"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
