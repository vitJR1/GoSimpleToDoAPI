package routes

import (
	_ "GoSimpleAPI/src/app/docs"
	"GoSimpleAPI/src/modules/todo"
	"GoSimpleAPI/src/modules/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupModules(router *gin.Engine, db *gorm.DB) {
	router.GET("/", HelloWorld)

	todo.InitTodoModule(router, db)
	users.InitUserModule(router, db)
}
