// /internal/src/users/repo/profile/Profile.repository.go
package profile

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}
