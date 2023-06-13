package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID       primitive.ObjectID `bson:"id"`
	Dish     *string            `json:"dish"`
	Macroes  *Macroes           `json:"macroes"`
	Calories *int64             `json:"calories"`
}

type Macroes struct {
	Protein int64 `json:"protein"`
	Fat     int64 `json:"fat"`
	Carbs   int64 `json:"carbs"`
}
