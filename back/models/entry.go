package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID       primitive.ObjectID `bson:"id"`
	Dish     *string            `json:"dish"`
	Macroes  *string            `json:"macroes"`
	Calories *int               `json:"calories"`
}
