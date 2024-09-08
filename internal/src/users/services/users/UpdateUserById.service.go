package users

import "go-rest/internal/src/users/domain"

func (u *UserService) UpdateUserByID(id string, user domain.UserUpdateDTO) (domain.UserResponseDTO, error) {
	_, err := u.GetUserByID(id)
	if err != nil {
		return domain.UserResponseDTO{}, err
	}
	updateUser, err := u.UserRepo.UpdateUserByID(id, user)
	if err != nil {
		return domain.UserResponseDTO{}, err
	}
	return updateUser, nil
}
