package profile

import "go-rest/internal/src/users/ports"

type profileService struct {
	profileRepo ports.ProfileRepository
}

func NewProfileService(profileRepo ports.ProfileRepository) *profileService {
	return &profileService{profileRepo: profileRepo}
}
