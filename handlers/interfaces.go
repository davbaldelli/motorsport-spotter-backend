package handlers

import (
	"motorsportspotter.backend/models"
	"net/http"
)

type ChampionshipsGateway interface {
	GETAllChampionships(http.ResponseWriter, *http.Request)
	POSTNewChampionship(http.ResponseWriter, *http.Request)
	UPDATEChampionship(http.ResponseWriter, *http.Request)
}

type TracksGateway interface {
	GETAllTracks(http.ResponseWriter, *http.Request)
	POSTNewTrack(http.ResponseWriter, *http.Request)
	UPDATETrack(http.ResponseWriter, *http.Request)
}

type EventsGateway interface {
	GETAllEvents(http.ResponseWriter, *http.Request)
	GETIncomingEvents(http.ResponseWriter, *http.Request)
	POSTNewEvent(http.ResponseWriter, *http.Request)
	UPDATEEvent(http.ResponseWriter, *http.Request)
}

type SessionsGateway interface {
	GETAllSessions(http.ResponseWriter, *http.Request)
	POSTNewSession(http.ResponseWriter, *http.Request)
	UPDATESession(http.ResponseWriter, *http.Request)
}

type NationsGateway interface {
	GETAllNations(http.ResponseWriter, *http.Request)
}

type UsersGateway interface {
	Login(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
}

type AuthorizationMiddleware interface {
	Authentication(next http.HandlerFunc) http.HandlerFunc
	Authorization(next http.HandlerFunc, allowedRoles []models.Role) http.HandlerFunc
}
