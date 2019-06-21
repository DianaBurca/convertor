package main

import (
	"github.com/DianaBurca/convertor/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	driver := gin.Default()

	driver.GET("/convert", utils.ConvertHandler)
	driver.GET("/.well-known/live", utils.Health)
	driver.GET("/.well-known/ready", utils.Health)

	driver.Run()
}
