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
type ProfileCreateDTO struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	Bio         string    `json:"bio" validate:"required"`
	AvatarURL   string    `json:"avatar_url" validate:"required"`
	DateOfBirth string    `json:"date_of_birth" validate:"required"`
}
type ProfileResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
