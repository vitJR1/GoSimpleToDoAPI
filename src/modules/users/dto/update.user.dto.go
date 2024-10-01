package dto

type UpdateUserDTO struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
