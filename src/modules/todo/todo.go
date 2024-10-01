package todo

import (
	"GoSimpleAPI/src/modules/todo/controller"
	"GoSimpleAPI/src/modules/todo/service"
	"gorm.io/gorm"
)

type Module struct {
	Controller *controller.TodoController
	Service    *service.TodoService
}

func NewTodoModule(db *gorm.DB) *Module {
	todoService := service.NewTodoService(db)
	todoController := controller.NewTodoController(todoService)

	return &Module{
		Controller: todoController,
		Service:    todoService,
	}
}
