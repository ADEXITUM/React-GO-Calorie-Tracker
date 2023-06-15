package api

import (
	"github.com/ADEXITUM/calorie-counter/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

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
