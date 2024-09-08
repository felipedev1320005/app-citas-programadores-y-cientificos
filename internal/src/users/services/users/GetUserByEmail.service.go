package users

import "go-rest/internal/src/users/domain"

func (u *UserService) GetUserByEmail(email string) (domain.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
