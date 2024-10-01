package dto

type CreateUserDTO struct {
	Name  string  `json:"name" binding:"required"`
	Email *string `json:"email"`
}
