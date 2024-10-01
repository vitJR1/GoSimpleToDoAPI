package todo

import (
	"gorm.io/gorm"
)

type Module struct {
	Controller *TodoController
	Service    *TodoService
}

func NewTodoModule(db *gorm.DB) *Module {
	todoService := NewTodoService(db)
	todoController := NewTodoController(todoService)

	return &Module{
		Controller: todoController,
		Service:    todoService,
	}
}
