package routes

import (
	"log"
	"os"

	"github.com/ADEXITUM/calorie-counter/models"
	"github.com/joho/godotenv"
)

func InitConfig() (*models.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("routes.InitConfig: #1\nError loading .env file\n%s\n\n", err)
		return nil, err
	}
	log.Println("Loaded config")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("routes.InitConfig: #2\nNo port provided")
	}

	mongoUri := os.Getenv("MONGO_URI")
	if port == "" {
		log.Fatal("routes.InitConfig: #2\nNo port provided")
	}

	cfg := LoadConfig(port, mongoUri)

	return cfg, nil
}

func LoadConfig(port, mongoUri string) *models.Config {
	return &models.Config{
		Port:     port,
		MongoURI: mongoUri,
	}
}
