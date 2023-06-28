package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExchangeRate struct {
	Cryptocurrency string             `bson:"cryptocurrency"`
	Fiat           string             `bson:"fiat"`
	Rate           float64            `bson:"rate"`
	CreatedAt      primitive.DateTime `bson:"createdAt"`
}
