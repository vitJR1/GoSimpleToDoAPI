package dto

type CreateTodoDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
}
