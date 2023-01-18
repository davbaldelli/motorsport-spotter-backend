package gateways

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type UsersGatewayImpl struct {
	Repo   repositories.UserRepository
	Secret string
}

func (u UsersGatewayImpl) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error in request body: %v ", err))
		return
	}

	authuser, err := u.Repo.Login(user)
	if err != nil {
		handlers.RespondError(w, http.StatusUnauthorized, err)
		return
	}

	validToken, err := GenerateJWT(authuser.Username, string(authuser.Role), u.Secret)
	if err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error generating token: %v ", err))
		return
	}

	token := models.Token{Username: authuser.Username, Role: string(authuser.Role), TokenString: validToken}

	handlers.RespondJSON(w, http.StatusAccepted, token)
}

func (u UsersGatewayImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error in request body: %v ", err))
		return
	}

	newUser, err := u.Repo.SignIn(user)
	if err != nil {
		handlers.RespondError(w, http.StatusUnauthorized, err)
		return
	}

	validToken, err := GenerateJWT(newUser.Username, string(newUser.Role), u.Secret)
	if err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error generating token: %v ", err))
		return
	}

	token := models.Token{Username: newUser.Username, Role: string(newUser.Role), TokenString: validToken}

	handlers.RespondJSON(w, http.StatusAccepted, token)
}
