package http

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type EventControllerImpl struct {
	Repo repositories.EventsRepository
}

func (c EventControllerImpl) GETAllEvents(writer http.ResponseWriter, _ *http.Request) {
	if tracks, err := c.Repo.GetAllEvents(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}
}

func (c EventControllerImpl) GETIncomingEvents(writer http.ResponseWriter, _ *http.Request) {
	if tracks, err := c.Repo.GetIncomingEvents(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}
}

func (c EventControllerImpl) POSTNewEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.InsertEvent(event); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, event)
}

func (c EventControllerImpl) UPDATEEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.UpdateEvent(event); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, event)
}
