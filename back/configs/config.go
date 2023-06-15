package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ADEXITUM/calorie-counter/pkg/entities"
	"github.com/joho/godotenv"
)

func InitConfig() (*entities.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("routes.InitConfig: #1\nError loading .env file\n%s\n\n", err)
		return nil, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("routes.InitConfig: #2\nNo port provided")
	}

	mongoUri := os.Getenv("MONGO_URI")
	if port == "" {
		log.Fatal("routes.InitConfig: #2\nNo port provided")
	}

	cfg := LoadConfig(port, mongoUri)

	fmt.Println("BOBO INA MO: " + mongoUri + port)

	return cfg, nil
}

func LoadConfig(port, mongoUri string) *entities.Config {
	return &entities.Config{
		Port:     port,
		MongoURI: mongoUri,
	}
}
