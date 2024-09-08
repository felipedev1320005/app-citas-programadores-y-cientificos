package profile

import (
	"go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/domain"
)

func (a *ProfileRepository) CreateProfile(profile domain.ProfileCreateDTO) (domain.ProfileResponseDTO, error) {
	newProfile := domain.Profile{
		UserID:      profile.UserID,
		Bio:         profile.Bio,
		AvatarURL:   profile.AvatarURL,
		DateOfBirth: profile.DateOfBirth,
	}
	createProfile := posgress.Db.Create(&newProfile)
	err := createProfile.Error
	if err != nil {
		return domain.ProfileResponseDTO{}, err
	}
	return domain.ProfileResponseDTO{
		ID:          newProfile.UserID,
		UserID:      newProfile.UserID,
		Bio:         newProfile.Bio,
		AvatarURL:   newProfile.AvatarURL,
		DateOfBirth: newProfile.DateOfBirth,
		CreatedAt:   newProfile.CreatedAt.String(),
		UpdatedAt:   newProfile.UpdatedAt.String(),
	}, nil
}
