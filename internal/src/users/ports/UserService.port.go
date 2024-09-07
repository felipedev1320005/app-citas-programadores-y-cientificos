// /internal/src/users/ports/UserService.port.go
package ports

import "go-rest/internal/src/users/domain"

type UserService interface {
	CreateUser(user domain.UserCreateDTO) (domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	ComparePasswords(hashedPassword string, password string) error
	DeleteUserByID(id string) error
}
