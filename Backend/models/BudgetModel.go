package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Budget struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty""`
	Amount int                `json:"amount,omitempty"`
	Tags   []string           `json:"tags,omitempty"`
	Date   primitive.DateTime `json:"date,omitempty"`
	User   *User              `json:"user,omitempty"`
}
