package http

import (
	"motorsportspotter.backend/repositories"
	"net/http"
)

type NationsControllerImpl struct {
	Repo repositories.NationsRepository
}

func (n NationsControllerImpl) GETAllNations(w http.ResponseWriter, r *http.Request) {
	if nations, err := n.Repo.GetAllNations(); err != nil {
		respondError(w, http.StatusInternalServerError, err)
	} else {
		respondJSON(w, http.StatusOK, nations)
	}
}
