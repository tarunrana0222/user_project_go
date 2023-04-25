package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	ClientID string             `json:"clientId"`
	Name     string             `json:"name"`
}
