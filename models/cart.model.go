package models

import "time"

type Cart struct {
	Id        string    `bson:"_id,omitempty"`
	UserId    string    `json:"user_id" bson:"user_id"`
	MenuId    string    `json:"menu_id" bson:"menu_id"`
	Quantity  int       `json:"quantity" bson:"quantity"`
	Status    string    `json:"status" bson:"status" default:"pending"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
