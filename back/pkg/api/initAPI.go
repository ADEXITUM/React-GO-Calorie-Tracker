package api

import (
	"log"

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
	log.Printf("\n\nGIN SUCCESSFULLY STARTED!\n\n")
	return nil
}
