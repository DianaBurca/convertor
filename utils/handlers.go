package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConvertHandler ...
func ConvertHandler(c *gin.Context) {
	params := c.Request.URL.Query()
	value := params["value"][0]

	var valueFloat float64
	var err error
	if valueFloat, err = strconv.ParseFloat(value, 64); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	fmt.Println(valueFloat)
	fmt.Println("ERR: ", err)

	output := map[string]interface{}{
		"cvalue": fmt.Sprintf("%.2f", valueFloat-273.15),
		"status": "200",
	}
	c.JSON(http.StatusOK, output)
}

// Health ...
func Health(c *gin.Context) {
	c.Status(http.StatusOK)
}
