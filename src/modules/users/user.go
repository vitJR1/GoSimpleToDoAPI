package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserModule(router *gin.Engine, db *gorm.DB) {
	userService := NewUserService(db)
	userController := NewUserController(userService)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("", userController.GetUserList)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.GET("/:id", userController.GetUserById)
	}
}
