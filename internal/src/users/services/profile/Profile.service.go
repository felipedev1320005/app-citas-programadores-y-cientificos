package profile

import (
	"go-rest/internal/src/users/ports/profile"
)

type profileService struct {
	profileRepo profile.ProfileRepository
}

func NewProfileService(profileRepo profile.ProfileRepository) *profileService {
	return &profileService{profileRepo: profileRepo}
}
