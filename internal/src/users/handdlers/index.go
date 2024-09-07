// /internal/src/users/handdlers/index.go
package handdlers

import (
	"encoding/json"
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	"net/http"
)

type UserHandler struct {
	UserRepository ports.UserRepository
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userBody domain.UserCreateDTO
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	newUser, err := u.UserRepository.CreateUser(userBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserRepository.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
// 	return []domain.User{}, nil
// }
// func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	return domain.User{}, nil
// }
// func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	return nil
// }
