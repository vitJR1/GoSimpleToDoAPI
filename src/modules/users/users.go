package users

import (
	"GoSimpleAPI/src/modules/users/controller"
	"GoSimpleAPI/src/modules/users/service"
	"gorm.io/gorm"
)

type Module struct {
	Controller *controller.UserController
	Service    *service.UserService
}

func NewUserModule(db *gorm.DB) *Module {
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	return &Module{
		Controller: userController,
		Service:    userService,
	}
}
