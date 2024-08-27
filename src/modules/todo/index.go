package todo

import (
	"GoSimpleAPI/src/modules/todo/controller"
	"GoSimpleAPI/src/modules/todo/service"
)

type Module struct {
	Controller *controller.TodoController
	Service    *service.TodoService
}

func NewTodoModule() *Module {
	todoService := service.NewTodoService()
	todoController := controller.NewTodoController(todoService)

	return &Module{
		Controller: todoController,
		Service:    todoService,
	}
}
