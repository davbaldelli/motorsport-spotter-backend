package gateways

import (
	"encoding/json"
	"fmt"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type ChampionshipGatewayImpl struct {
	Repo repositories.ChampionshipsRepository
}

func (c ChampionshipGatewayImpl) GETAllChampionships(writer http.ResponseWriter, _ *http.Request) {
	if champs, err := c.Repo.GetAllChampionships(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, champs)
	}
}

func (c ChampionshipGatewayImpl) POSTNewChampionship(w http.ResponseWriter, r *http.Request) {
	championship := models.Championship{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&championship); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.InsertChampionship(championship); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, championship)
}

func (c ChampionshipGatewayImpl) UPDATEChampionship(w http.ResponseWriter, r *http.Request) {
	championship := models.Championship{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&championship); err != nil {
		handlers.RespondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.Repo.UpdateChampionship(championship); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, fmt.Errorf("error adding new entity: %v", err))
		return
	}

	handlers.RespondJSON(w, http.StatusOK, championship)
}
