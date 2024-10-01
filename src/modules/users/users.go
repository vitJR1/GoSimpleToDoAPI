package users

import (
	"gorm.io/gorm"
)

type Module struct {
	Controller *UserController
	Service    *UserService
}

func NewUserModule(db *gorm.DB) *Module {
	userService := NewUserService(db)
	userController := NewUserController(userService)

	return &Module{
		Controller: userController,
		Service:    userService,
	}
}
