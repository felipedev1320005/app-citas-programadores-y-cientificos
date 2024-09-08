// /internal/src/users/ports/profile/ProfileRepository.port.go
package profile

import "go-rest/internal/src/users/domain"

type ProfileRepository interface {
	CreateProfile(profile domain.ProfileCreateDTO) (domain.ProfileResponseDTO, error)
}
