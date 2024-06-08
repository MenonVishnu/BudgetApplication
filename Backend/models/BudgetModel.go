package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Budget struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Amount int                `json:"amount,omitempty" validate:"required"`
	Tags   []string           `json:"tags,omitempty" validate:"required"` 
	Date   primitive.DateTime `json:"date,omitempty" validate:"required"`
	User   *User              `json:"user,omitempty" validate:"required"`
}
