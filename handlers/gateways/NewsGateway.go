package gateways

import (
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/repositories"
	"net/http"
)

type NewsGatewayImpl struct {
	Repo repositories.NewsRepository
}

func (n NewsGatewayImpl) GETAllNews(writer http.ResponseWriter, _ *http.Request) {
	if news, err := n.Repo.GetAllNews(); err != nil {
		handlers.RespondError(writer, http.StatusInternalServerError, err)
	} else {
		handlers.RespondJSON(writer, http.StatusOK, news)
	}
}
