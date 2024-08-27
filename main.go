package main

import (
	"GoSimpleAPI/src/app/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GoSimpleAPI
// @version         1.0
// @description     This is a simple API.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.SetupRoutes(r)

	r.Run()
}
