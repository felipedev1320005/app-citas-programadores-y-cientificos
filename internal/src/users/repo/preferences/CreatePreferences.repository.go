package preferences

import (
	"go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/domain"
)

func (a *PreferencesRepository) CreatePreferences(preferences domain.PreferencesCreateDTO) (domain.PreferencesResponseDTO, error) {
	createPreferences := posgress.Db.Create(&preferences)
	err := createPreferences.Error
	if err != nil {
		return domain.PreferencesResponseDTO{}, err
	}
	preferencesResponse := domain.PreferencesResponseDTO{
		ReceiveEmails:     preferences.ReceiveEmails,
		ShowOnlineStatus:  preferences.ShowOnlineStatus,
		PreferredLanguage: preferences.PreferredLanguage,
		UserID:            preferences.UserID,
	}
	return preferencesResponse, nil
}
