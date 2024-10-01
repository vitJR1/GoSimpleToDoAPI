package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserModule(router *gin.Engine, db *gorm.DB) {
	userController := NewUserModule(db).Controller

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("", userController.GetUserList)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.GET("/:id", userController.GetUserById)
	}
}
