package routes

import (
	"GoSimpleAPI/src/modules/todo/controller"
	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine, todoController *controller.TodoController) {
	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("", todoController.GetToDoList)
		todoRoutes.POST("", todoController.CreateTodo)
		todoRoutes.PUT("/:id", todoController.UpdateTodo)
		todoRoutes.GET("/:id", todoController.GetToDoById)
	}
}
