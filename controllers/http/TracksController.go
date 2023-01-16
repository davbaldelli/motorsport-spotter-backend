package http

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type TracksControllerImpl struct {
	Repo repositories.TracksRepository
}

func (t TracksControllerImpl) POSTNewTrack(w http.ResponseWriter, r *http.Request) {
	track := models.Track{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&track); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.Repo.InsertTrack(track); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, track)
}

func (t TracksControllerImpl) UPDATETrack(w http.ResponseWriter, r *http.Request) {
	track := models.Track{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&track); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.Repo.UpdateTrack(track); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, track)
}

func (t TracksControllerImpl) GETAllTracks(writer http.ResponseWriter, request *http.Request) {
	if tracks, err := t.Repo.GetAllTracks(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}
}
