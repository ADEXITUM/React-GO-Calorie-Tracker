package api

import (
	"github.com/ADEXITUM/calorie-counter/pkg/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ADEXITUM/calorie-counter/docs"
)

func initHandlers(engine *gin.Engine) {
	engine.GET("/entries/", handlers.GetEntries)
	engine.GET("/entries/:id/", handlers.GetEntryById)
	engine.POST("/entries/create/", handlers.AddEntry)
	engine.PUT("/entries/update/:id", handlers.UpdateEntry)
	engine.DELETE("/entries/delete/all/", handlers.DeleteAll)
	engine.DELETE("/entries/delete/:id/", handlers.DeleteEntry)

	engine.POST("/calculate/calories/", handlers.CalculateCalories)
	engine.POST("/calculate/massindex/", handlers.CalculateMassIndex)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
