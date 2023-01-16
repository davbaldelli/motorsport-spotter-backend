package http

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type SessionsControllerImpl struct {
	Repo repositories.SessionsRepository
}

func (s SessionsControllerImpl) GETAllSessions(writer http.ResponseWriter, request *http.Request) {
	if sessions, err := s.Repo.GetAllSessions(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, sessions)
	}
}

func (s SessionsControllerImpl) POSTNewSession(w http.ResponseWriter, r *http.Request) {
	session := models.Session{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&session); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Repo.InsertSession(session); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, session)
}

func (s SessionsControllerImpl) UPDATESession(w http.ResponseWriter, r *http.Request) {
	session := models.Session{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&session); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Repo.UpdateSession(session); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, session)
}
