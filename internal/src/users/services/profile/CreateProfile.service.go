package profile

import "go-rest/internal/src/users/domain"

func (a *profileService) CreateProfile(dto domain.ProfileCreateDTO) (*domain.ProfileResponseDTO, error) {
	// Create Profile
	newProfile, err := a.profileRepo.CreateProfile(dto)
	if err != nil {
		return nil, err
	}
	return &newProfile, nil
}
