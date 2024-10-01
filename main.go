package main

import (
	"GoSimpleAPI/src/app/config"
	"GoSimpleAPI/src/app/database"
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
	err := config.Init()
	if err != nil {
		return
	}

	err = database.Connect(config.CFG)

	if err != nil {
		return
	}

	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.SetupModules(r, database.DB)

	err = r.Run()
	if err != nil {
		return
	}
}
