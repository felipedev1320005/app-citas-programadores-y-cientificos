package users

import "go-rest/internal/src/users/domain"

func (u *UserService) GetUsers() ([]domain.User, error) {
	user, err := u.UserRepo.GetUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return user, nil
}
