package http

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Web struct {
	ChampCtrl    ChampionshipsController
	TracksCtrl   TracksController
	EventCtrl    EventController
	SessionsCtrl SessionsController
	NewsCtrl     NewsController
}

func (w Web) Listen() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/championships", w.ChampCtrl.GETAllChampionships).Methods("GET")
	router.HandleFunc("/api/championships/add", w.ChampCtrl.POSTNewChampionship).Methods("POST")
	router.HandleFunc("/api/championships/update", w.ChampCtrl.UPDATEChampionship).Methods("POST")

	router.HandleFunc("/api/tracks", w.TracksCtrl.GETAllTracks).Methods("GET")
	router.HandleFunc("/api/tracks/add", w.TracksCtrl.POSTNewTrack).Methods("POST")
	router.HandleFunc("/api/tracks/update", w.TracksCtrl.UPDATETrack).Methods("POST")

	router.HandleFunc("/api/events", w.EventCtrl.GETAllEvents).Methods("GET")
	router.HandleFunc("/api/events/incoming", w.EventCtrl.GETIncomingEvents).Methods("GET")
	router.HandleFunc("/api/events/add", w.EventCtrl.POSTNewEvent).Methods("POST")
	router.HandleFunc("/api/events/update", w.EventCtrl.UPDATEEvent).Methods("POST")

	router.HandleFunc("/api/sessions", w.SessionsCtrl.GETAllSessions).Methods("GET")
	router.HandleFunc("/api/sessions/add", w.SessionsCtrl.POSTNewSession).Methods("POST")
	router.HandleFunc("/api/sessions/update", w.SessionsCtrl.UPDATESession).Methods("POST")

	router.HandleFunc("/api/news", w.NewsCtrl.GETAllNews).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":7151", handler))
}
