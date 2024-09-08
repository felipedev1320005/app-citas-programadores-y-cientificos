package adapters

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports/profile"
)

type profileToUserAdapter struct {
	ProgileService profile.ProfileService
}

func NewProfileToUserAdapter(profileService profile.ProfileService) *profileToUserAdapter {
	return &profileToUserAdapter{ProgileService: profileService}
}
func (p *profileToUserAdapter) CreateProfile(profileDto domain.ProfileCreateDTO) (*domain.ProfileResponseDTO, error) {
	newProfile, err := p.ProgileService.CreateProfile(profileDto)
	if err != nil {
		return nil, err
	}
	return newProfile, nil
}
