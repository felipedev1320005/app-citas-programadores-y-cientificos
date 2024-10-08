// /internal/src/users/repo/User.repository.go
package repo

import (
	"go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
func (u *UserRepository) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	newUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword), // Guardar la contraseña encriptada
	}

	// Guardar el nuevo usuario en la base de datos
	createUser := posgress.Db.Create(&newUser)
	err = createUser.Error
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

func (u *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	getUsers := posgress.Db.Find(&users)
	err := getUsers.Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
func (u *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	getUser := posgress.Db.Where("email = ?", email).First(&user)
	err := getUser.Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u *UserRepository) DeleteUserByID(id string) error {
	deleteUser := posgress.Db.Where("id = ?", id).Delete(&domain.User{})
	err := deleteUser.Error
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepository) GetUserByID(id string) (domain.User, error) {
	var user domain.User
	getUser := posgress.Db.Where("id = ?", id).First(&user)
	err := getUser.Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u *UserRepository) UpdateUserByID(id string, user domain.UserUpdateDTO) (domain.UserResponseDTO, error) {
	updateUser := posgress.Db.Model(&domain.User{}).Where("id = ?", id).Updates(user)
	err := updateUser.Error
	if err != nil {
		return domain.UserResponseDTO{}, err
	}
	userDto := domain.UserResponseDTO{
		Name:  user.Name,
		Email: user.Email,
	}
	return userDto, nil
}
