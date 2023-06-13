package routes

import (
	"log"

	"github.com/ADEXITUM/calorie-counter/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, cfg *models.Config) error {
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create/", AddEntry)
	router.POST("/calculate/calories/", CalculateCalories)
	router.POST("/calculate/massindex/", CalculateMassIndex)

	router.GET("/entries/", GetEntries)
	router.GET("/entries/:id/", GetEntryById)
	router.GET("/dish/:dish/", GetEntriesByDish)

	router.PUT("/entry/update/:id", UpdateEntry)

	router.DELETE("/entry/delete/all/", DeleteAll)
	router.DELETE("/entry/delete/:id/", DeleteEntry)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Printf("routes.InitRoutes: #1\nError while running router:\n%s\n\n", err)
		return err
	}
	return nil
}
