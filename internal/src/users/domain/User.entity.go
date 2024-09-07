// /internal/src/users/domain/User.entity.go
package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string       `gorm:"not null"`
	Email       string       `gorm:"unique;not null"`
	Password    string       `gorm:"not null"`
	Profile     *Profile     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Preferences *Preferences `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type UserCreateDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
