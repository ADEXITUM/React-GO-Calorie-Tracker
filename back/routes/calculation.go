package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ADEXITUM/calorie-counter/models"
	"github.com/gin-gonic/gin"
)

func CalculateCalories(c *gin.Context) {
	var calculation models.Calculation
	var calories float32

	err := c.BindJSON(&calculation)
	if err != nil {
		log.Println(err)
	}

	switch calculation.Gender {
	case "Male":
		calories = 88.36 + (13.4 * calculation.Weight) + (4.8 * calculation.Height) - (5.7 * calculation.Age)
	case "Female":
		calories = 447.6 + (9.2 * calculation.Weight) + (3.1 * calculation.Height) - (5.7 * calculation.Age)
	default:
		c.JSON(http.StatusBadRequest, "Incorrect gender")
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%.2f", calories))
}

func CalculateMassIndex(c *gin.Context) {
	var calculation models.Calculation
	var massIndex float32

	err := c.BindJSON(&calculation)
	if err != nil {
		log.Println(err)
	}

	massIndex = calculation.Weight / ((calculation.Height / 100) * (calculation.Height / 100))

	c.JSON(http.StatusOK, fmt.Sprintf("%.2f", massIndex))
}
