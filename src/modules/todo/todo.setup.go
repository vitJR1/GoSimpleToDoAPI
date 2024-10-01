package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTodoModule(router *gin.Engine, db *gorm.DB) {
	todoController := NewTodoModule(db).Controller
	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("", todoController.GetToDoList)
		todoRoutes.POST("", todoController.CreateTodo)
		todoRoutes.PUT("/:id", todoController.UpdateTodo)
		todoRoutes.GET("/:id", todoController.GetToDoById)
	}
}
