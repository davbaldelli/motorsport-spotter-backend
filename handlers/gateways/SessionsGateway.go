package gateways

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type SessionsGatewayImpl struct {
	Repo repositories.SessionsRepository
}

func (s SessionsGatewayImpl) GETAllSessions(writer http.ResponseWriter, request *http.Request) {
	if sessions, err := s.Repo.GetAllSessions(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, sessions)
	}
}

func (s SessionsGatewayImpl) POSTNewSession(w http.ResponseWriter, r *http.Request) {
	session := models.Session{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&session); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Repo.InsertSession(session); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, session)
}

func (s SessionsGatewayImpl) UPDATESession(w http.ResponseWriter, r *http.Request) {
	session := models.Session{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&session); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Repo.UpdateSession(session); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, session)
}
