package http

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type ChampionshipControllerImpl struct {
	Repo repositories.ChampionshipsRepository
}

func (c ChampionshipControllerImpl) GETAllChampionships(writer http.ResponseWriter, _ *http.Request) {
	if champs, err := c.Repo.GetAllChampionships(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}
}

func (c ChampionshipControllerImpl) POSTNewChampionship(w http.ResponseWriter, r *http.Request) {
	championship := models.Championship{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&championship); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.InsertChampionship(championship); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, championship)
}

func (c ChampionshipControllerImpl) UPDATEChampionship(w http.ResponseWriter, r *http.Request) {
	championship := models.Championship{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&championship); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.UpdateChampionship(championship); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	respondJSON(w, http.StatusOK, championship)
}
