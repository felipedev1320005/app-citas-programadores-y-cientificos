// /internal/src/users/domain/Preferences.entity.go
package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Preferences struct {
	gorm.Model
	UserID            uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	ReceiveEmails     bool      `gorm:"not null;default:true"`
	ShowOnlineStatus  bool      `gorm:"not null;default:true"`
	PreferredLanguage string    `gorm:"type:varchar(50)"`
	User              User      `gorm:"foreignKey:UserID"`
}
