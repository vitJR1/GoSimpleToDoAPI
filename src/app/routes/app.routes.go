package routes

import (
	_ "GoSimpleAPI/src/app/docs"
	"GoSimpleAPI/src/modules/todo"
	"GoSimpleAPI/src/modules/todo/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", HelloWorld)

	todoModule := todo.NewTodoModule()
	routes.SetupTodoRoutes(router, todoModule.Controller)
}
