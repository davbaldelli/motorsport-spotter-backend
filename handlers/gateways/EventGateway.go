package gateways

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type EventGatewayImpl struct {
	Repo repositories.EventsRepository
}

func (c EventGatewayImpl) GETAllEvents(writer http.ResponseWriter, _ *http.Request) {
	if tracks, err := c.Repo.GetAllEvents(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, tracks)
	}
}

func (c EventGatewayImpl) GETIncomingEvents(writer http.ResponseWriter, _ *http.Request) {
	if tracks, err := c.Repo.GetIncomingEvents(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, tracks)
	}
}

func (c EventGatewayImpl) POSTNewEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.InsertEvent(event); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, event)
}

func (c EventGatewayImpl) UPDATEEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.UpdateEvent(event); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, event)
}
