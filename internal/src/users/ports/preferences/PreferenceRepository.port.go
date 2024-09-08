package preferences

import "go-rest/internal/src/users/domain"

type PreferencesRepository interface {
	CreatePreferences(preferences domain.PreferencesCreateDTO) (domain.PreferencesResponseDTO, error)
}
