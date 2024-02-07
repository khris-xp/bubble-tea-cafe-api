package models

import (
	"time"
)

type Category struct {
	Id        string    `bson:"_id,omitempty"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
