package main

import (
	"log"

	config "github.com/ADEXITUM/calorie-counter/configs"
	api "github.com/ADEXITUM/calorie-counter/pkg/api"
	"github.com/gin-gonic/gin"
)

// @title Calorie Tracker
// @version 1.0.0
// @decription Back-end for Calorie Tracker Application

// @host localhost:8080
// @BasePath /

func main() {
	engine := gin.New()
	log.Println("Created a gin instance")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Printf("FATAL ERROR: couldn't init config")
	}
	log.Println("Loaded config")

	err = api.InitAPI(engine, cfg)
	if err != nil {
		log.Printf("FATAL ERROR: couldn't init routes")
	}
	log.Println("Initialized API")
}
