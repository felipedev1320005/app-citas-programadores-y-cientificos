package users

import "go-rest/internal/src/users/domain"

func (u *UserService) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser, err := u.UserRepo.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
