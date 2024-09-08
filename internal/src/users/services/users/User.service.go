package users

import (
	"go-rest/internal/src/users/ports"
	"go-rest/internal/src/users/ports/profile"
)

type UserService struct {
	UserRepo       ports.UserRepository
	profileService profile.ProfileToUserAdapter
}

func NewUserService(userRepo ports.UserRepository, profileService profile.ProfileToUserAdapter) *UserService {
	return &UserService{
		UserRepo:       userRepo,
		profileService: profileService,
	}
}
