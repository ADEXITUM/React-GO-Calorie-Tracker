package main

import (
	"log"

	"github.com/ADEXITUM/calorie-counter/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	log.Println("Created a gin instance")

	cfg, err := routes.InitConfig()
	if err != nil {
		log.Printf("FATAL ERROR: couldn't init config")
	}
	log.Println("Loaded config")

	err = routes.InitRoutes(router, cfg)
	if err != nil {
		log.Printf("FATAL ERROR: couldn't init routes")
	}
	log.Println("Set up gin routes")
}
