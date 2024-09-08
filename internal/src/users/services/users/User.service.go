// /internal/src/users/services/users/User.service.go
package users

import (
	"go-rest/internal/src/users/ports"
	"go-rest/internal/src/users/ports/profile"
)

type UserService struct {
	UserRepo       ports.UserRepository
	profileService profile.ProfileToUserAdapter
}

// Constructor de UserService, implementa el patron de diseño Singleton y inyección de dependencias
func NewUserService(userRepo ports.UserRepository, profileService profile.ProfileToUserAdapter) *UserService {
	return &UserService{
		UserRepo:       userRepo,
		profileService: profileService,
	}
}
