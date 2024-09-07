package ports

import "go-rest/internal/src/users/domain"

type UserService interface {
	CreateUser(user domain.UserCreateDTO) (domain.User, error)
	GetUsers() ([]domain.User, error)
}
