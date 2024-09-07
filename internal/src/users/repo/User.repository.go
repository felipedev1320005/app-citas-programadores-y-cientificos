// /internal/src/users/repo/User.repository.go
package repo

import (
	"go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/domain"
)

type UserRepository struct {
}

// func (u *UserRepository) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
// 	newUser, err := u.UserService.CreateUser(user)
// 	if err != nil {
// 		return domain.User{}, err
// 	}
// 	return newUser, nil
// }
// func (u *UserRepository) GetUsers() ([]domain.User, error) {
// 	user, err := u.UserService.GetUsers()
// 	if err != nil {
// 		return []domain.User{}, err
// 	}
// 	return user, nil
// }

func (u *UserRepository) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	createUser := posgress.Db.Create(&newUser)
	err := createUser.Error
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
