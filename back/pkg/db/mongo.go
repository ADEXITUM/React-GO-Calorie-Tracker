package db

import (
	"context"
	"log"
	"time"

	config "github.com/ADEXITUM/calorie-counter/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Printf("routes.DBinstance: #1\nError setting up DB\n%s\n\n", err)
		return nil
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Printf("routes.DBinstance: #2\nError setting up DB\n%s\n\n", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("routes.DBinstance: #3\nError setting up DB\n%s\n\n", err)
		return nil
	}

	log.Println("Connected to MongoDB")

	return client
}

func OpenCollection(client *mongo.Client, name string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(name)

	return collection
}
