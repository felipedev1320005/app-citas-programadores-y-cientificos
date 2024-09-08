package users

func (u *UserService) DeleteUserByID(id string) error {
	err := u.UserRepo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
