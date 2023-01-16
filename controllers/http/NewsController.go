package http

import (
	"motorsportspotter.backend/repositories"
	"net/http"
)

type NewsControllerImpl struct {
	Repo repositories.NewsRepository
}

func (n NewsControllerImpl) GETAllNews(writer http.ResponseWriter, _ *http.Request) {
	if news, err := n.Repo.GetAllNews(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, news)
	}
}
