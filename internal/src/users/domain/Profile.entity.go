// /internal/src/users/domain/Profile.entity.go
package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID      uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Bio         string    `gorm:"type:text"`
	AvatarURL   string    `gorm:"type:text"`
	DateOfBirth string    `gorm:"type:date"`
	User        User      `gorm:"foreignKey:UserID"`
}
