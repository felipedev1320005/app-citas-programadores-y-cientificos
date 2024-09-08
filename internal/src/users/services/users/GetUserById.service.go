package users

import "go-rest/internal/src/users/domain"

func (u *UserService) GetUserByID(id string) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
