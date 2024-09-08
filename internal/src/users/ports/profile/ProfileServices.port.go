// /internal/src/users/ports/profile/ProfileServices.port.go
package profile

import "go-rest/internal/src/users/domain"

type ProfileService interface {
	CreateProfile(profileDto domain.ProfileCreateDTO) (*domain.ProfileResponseDTO, error)
}
