package controllers

import "net/http"

type ChampionshipsController interface {
	GETAllChampionships(http.ResponseWriter, *http.Request)
	POSTNewChampionship(http.ResponseWriter, *http.Request)
	UPDATEChampionship(http.ResponseWriter, *http.Request)
}

type TracksController interface {
	GETAllTracks(http.ResponseWriter, *http.Request)
	POSTNewTrack(http.ResponseWriter, *http.Request)
	UPDATETrack(http.ResponseWriter, *http.Request)
}

type EventController interface {
	GETAllEvents(http.ResponseWriter, *http.Request)
	GETIncomingEvents(http.ResponseWriter, *http.Request)
	POSTNewEvent(http.ResponseWriter, *http.Request)
	UPDATEEvent(http.ResponseWriter, *http.Request)
}

type SessionsController interface {
	GETAllSessions(http.ResponseWriter, *http.Request)
	POSTNewSession(http.ResponseWriter, *http.Request)
	UPDATESession(http.ResponseWriter, *http.Request)
}

type NewsController interface {
	GETAllNews(http.ResponseWriter, *http.Request)
}

type NationsController interface {
	GETAllNations(http.ResponseWriter, *http.Request)
}
