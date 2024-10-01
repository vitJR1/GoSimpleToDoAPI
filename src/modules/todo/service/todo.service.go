package service

import (
	"GoSimpleAPI/src/modules/todo/dto"
	"GoSimpleAPI/src/modules/todo/entity"
	"errors"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type TodoService struct {
	todos []entity.Todo
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		todos: []entity.Todo{},
	}
}

func (s *TodoService) CreateTodo(data dto.CreateTodoDTO) entity.Todo {
	newID := len(s.todos) + 1

	newTodo := entity.Todo{
		ID:          uint(newID),
		Title:       data.Title,
		Description: *data.Description,
	}

	s.todos = append(s.todos, newTodo)

	return newTodo
}

func (s *TodoService) UpdateTodo(id int, data dto.UpdateTodoDTO) (*entity.Todo, error) {
	for i, todo := range s.todos {
		if todo.ID == uint(id) {
			// Обновляем поля задачи
			if data.Title != nil {
				s.todos[i].Title = *data.Title
			}
			if data.Description != nil {
				s.todos[i].Description = *data.Description
			}
			return &s.todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func (s *TodoService) GetTodoList(search string) []entity.Todo {
	var filteredTodos []entity.Todo

	// Filter todos based on the search string
	for _, todo := range s.todos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(search)) {
			filteredTodos = append(filteredTodos, todo)
		}
	}

	// Sort the filtered todos by title
	sort.Slice(filteredTodos, func(i, j int) bool {
		return filteredTodos[i].Title < filteredTodos[j].Title
	})

	return filteredTodos
}

func (s *TodoService) GetTodoById(id int) *entity.Todo {
	for _, todo := range s.todos {
		if todo.ID == uint(id) {
			return &todo
		}
	}
	return nil
}
