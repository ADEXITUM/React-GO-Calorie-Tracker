package routes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	MongoURI := "mongodb://localhost:27017/caloriesdb"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Printf("routes.DBinstance: #1\nError setting up DB\n%s\n\n", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("routes.DBinstance: #1\nError setting up DB\n%s\n\n", err)
		return nil
	}

	log.Println("Connected to MongoDB")

	return client
}

func openCollection(client *mongo.Client, name string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(name)

	return collection
}
