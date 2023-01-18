package gateways

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type TracksGatewayImpl struct {
	Repo repositories.TracksRepository
}

func (t TracksGatewayImpl) POSTNewTrack(w http.ResponseWriter, r *http.Request) {
	track := models.Track{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&track); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.Repo.InsertTrack(track); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, track)
}

func (t TracksGatewayImpl) UPDATETrack(w http.ResponseWriter, r *http.Request) {
	track := models.Track{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&track); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.Repo.UpdateTrack(track); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, track)
}

func (t TracksGatewayImpl) GETAllTracks(writer http.ResponseWriter, request *http.Request) {
	if tracks, err := t.Repo.GetAllTracks(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, tracks)
	}
}
