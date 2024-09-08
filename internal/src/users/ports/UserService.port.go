// /internal/src/users/ports/UserService.port.go
package ports

import "go-rest/internal/src/users/domain"

type UserService interface {
	CreateUser(user domain.UserCreateDTO) (domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	ComparePasswords(hashedPassword string, password string) error
	DeleteUserByID(id string) error
	GetUserByID(id string) (domain.User, error)
	FormateUser(user domain.User) domain.UserResponseDTO
	FormateUsers(users []domain.User) []domain.UserResponseDTO
}
