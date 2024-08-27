package dto

type UpdateTodoDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
