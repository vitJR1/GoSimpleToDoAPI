package users

import (
	"GoSimpleAPI/src/modules/users/dto"
	"GoSimpleAPI/src/modules/users/entity"
	"errors"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type UserService struct {
	users []entity.User
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		users: []entity.User{},
	}
}

func (s *UserService) CreateUser(data dto.CreateUserDTO) entity.User {
	newID := len(s.users) + 1

	newUser := entity.User{
		ID:    uint(newID),
		Name:  data.Name,
		Email: *data.Email,
	}

	s.users = append(s.users, newUser)

	return newUser
}

func (s *UserService) UpdateUser(id int, data dto.UpdateUserDTO) (*entity.User, error) {
	for i, user := range s.users {
		if user.ID == uint(id) {
			// Обновляем поля задачи
			if data.Name != nil {
				s.users[i].Name = *data.Name
			}
			if data.Email != nil {
				s.users[i].Email = *data.Email
			}
			return &s.users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *UserService) GetUserList(search string) []entity.User {
	var filteredUsers []entity.User

	// Filter users based on the search string
	for _, user := range s.users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(search)) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// Sort the filtered users by title
	sort.Slice(filteredUsers, func(i, j int) bool {
		return filteredUsers[i].Name < filteredUsers[j].Name
	})

	return filteredUsers
}

func (s *UserService) GetUserById(id int) *entity.User {
	for _, user := range s.users {
		if user.ID == uint(id) {
			return &user
		}
	}
	return nil
}
