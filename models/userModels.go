package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID string             `json:"userId" bson:"userId" validate:"required"`
	Client string             `json:"client" validate:"required"`
	Name   string             `json:"name" validate:"required"`
	Os     string             `json:"os"`
}
