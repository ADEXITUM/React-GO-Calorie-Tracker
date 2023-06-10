package main

import (
	"log"
	"os"

	"github.com/ADEXITUM/calorie-counter/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("#1: Error loading .env file\n%s", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("#2: Error loading port")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries/", routes.GetEntries)
	router.GET("/entries/:id/", routes.GetEntryById)

	router.PUT("/entry/update/:id/", routes.UpdateEntry)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)
}
