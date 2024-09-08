package users

import "go-rest/internal/src/users/domain"

func (u *UserService) FormateUser(user domain.User) domain.UserResponseDTO {
	return domain.UserResponseDTO{
		Name:        user.Name,
		Email:       user.Email,
		Profile:     user.Profile,
		Preferences: user.Preferences,
	}
}
func (u *UserService) FormateUsers(users []domain.User) []domain.UserResponseDTO {
	var response []domain.UserResponseDTO
	for _, user := range users {
		response = append(response, domain.UserResponseDTO{
			Name:        user.Name,
			Email:       user.Email,
			Profile:     user.Profile,
			Preferences: user.Preferences,
		})
	}
	return response
}
