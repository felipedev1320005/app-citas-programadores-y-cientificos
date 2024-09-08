// /internal/src/users/ports/profile/ProfileToUserAdapter.port.go
package profile

import "go-rest/internal/src/users/domain"

type ProfileToUserAdapter interface {
	CreateProfile(profileDto domain.ProfileCreateDTO) (*domain.ProfileResponseDTO, error)
}
