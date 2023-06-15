package api

import (
	"log"

	"github.com/ADEXITUM/calorie-counter/pkg/api/handlers"
	"github.com/ADEXITUM/calorie-counter/pkg/entities"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitAPI(engine *gin.Engine, cfg *entities.Config) error {
	engine.Use(gin.Logger())
	engine.Use(cors.Default())

	initHandlers(engine)
	log.Println("Initialized handlers")

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		log.Printf("routes.InitRoutes: #1\nError while running router:\n%s\n\n", err)
		return err
	}
	log.Println("\n\nGIN SUCCESSFULLY STARTED!\n\n")
	return nil
}

func initHandlers(engine *gin.Engine) {
	engine.GET("/entries/", handlers.GetEntries)
	engine.GET("/entries/:id/", handlers.GetEntryById)
	engine.GET("/dish/:dish/", handlers.GetEntriesByDish)

	engine.POST("/entry/create/", handlers.AddEntry)
	engine.POST("/calculate/calories/", handlers.CalculateCalories)
	engine.POST("/calculate/massindex/", handlers.CalculateMassIndex)

	engine.PUT("/entry/update/:id", handlers.UpdateEntry)

	engine.DELETE("/entry/delete/all/", handlers.DeleteAll)
	engine.DELETE("/entry/delete/:id/", handlers.DeleteEntry)
}
