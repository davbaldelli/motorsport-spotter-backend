package gateways

import (
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type NationsGatewayImpl struct {
	Repo repositories.NationsRepository
}

func (n NationsGatewayImpl) GETAllNations(w http.ResponseWriter, r *http.Request) {
	if nations, err := n.Repo.GetAllNations(); err != nil {
		handlers.RespondError(w, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(w, http.StatusOK, nations)
	}
}
