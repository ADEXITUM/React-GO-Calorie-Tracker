package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ADEXITUM/calorie-counter/pkg/entities"

	"github.com/gin-gonic/gin"
)

// @Summary Calculate Calories
// @Tags calculation
// @Description Calculate calories based on Harris-Benedicts formula
// @ID calculate-calories
// @Accept json
// @Produce json
// @Param input body entities.Calculation true "calculation info"
// @Success 200 {number} calories
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /calculate/calories [post]
func CalculateCalories(c *gin.Context) {
	var calculation entities.Calculation
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

// @Summary Calculate Body Mass Index
// @Tags calculation
// @Description Calculate Body Mass Index based on Adolf Ketles formula
// @ID calculate-massidex
// @Accept json
// @Produce json
// @Param input body entities.Calculation true "calculation info"
// @Success 200 {number} massIndex
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /calculate/massindex [post]
func CalculateMassIndex(c *gin.Context) {
	var calculation entities.Calculation
	var massIndex float32

	err := c.BindJSON(&calculation)
	if err != nil {
		log.Println(err)
	}

	massIndex = calculation.Weight / ((calculation.Height / 100) * (calculation.Height / 100))

	c.JSON(http.StatusOK, fmt.Sprintf("%.2f", massIndex))
}
