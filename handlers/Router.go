package handlers

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"motorsportspotter.backend/models"
	"net/http"
	"sync"
)

type Router struct {
	ChampGate    ChampionshipsGateway
	TracksGate   TracksGateway
	EventGate    EventsGateway
	SessionsGate SessionsGateway
	NationGate   NationsGateway
	UserGate     UsersGateway
	AuthMidl     AuthorizationMiddleware
}

func (w Router) Listen() {

	adminOnly := []models.Role{models.Admin}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/championships", w.AuthMidl.Authentication(w.ChampGate.GETAllChampionships)).Methods("GET")
	router.HandleFunc("/api/championships/add", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.ChampGate.POSTNewChampionship, adminOnly))).Methods("POST")
	router.HandleFunc("/api/championships/update", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.ChampGate.UPDATEChampionship, adminOnly))).Methods("POST")

	router.HandleFunc("/api/tracks", w.AuthMidl.Authentication(w.TracksGate.GETAllTracks)).Methods("GET")
	router.HandleFunc("/api/tracks/add", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.TracksGate.POSTNewTrack, adminOnly))).Methods("POST")
	router.HandleFunc("/api/tracks/update", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.TracksGate.UPDATETrack, adminOnly))).Methods("POST")

	router.HandleFunc("/api/events", w.AuthMidl.Authentication(w.EventGate.GETAllEvents)).Methods("GET")
	router.HandleFunc("/api/events/incoming", w.AuthMidl.Authentication(w.EventGate.GETIncomingEvents)).Methods("GET")
	router.HandleFunc("/api/events/add", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.EventGate.POSTNewEvent, adminOnly))).Methods("POST")
	router.HandleFunc("/api/events/update", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.EventGate.UPDATEEvent, adminOnly))).Methods("POST")

	router.HandleFunc("/api/sessions", w.AuthMidl.Authentication(w.SessionsGate.GETAllSessions)).Methods("GET")
	router.HandleFunc("/api/sessions/add", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.SessionsGate.POSTNewSession, adminOnly))).Methods("POST")
	router.HandleFunc("/api/sessions/update", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.SessionsGate.UPDATESession, adminOnly))).Methods("POST")

	router.HandleFunc("/api/login", w.UserGate.Login).Methods("POST")
	router.HandleFunc("/api/signin", w.AuthMidl.Authentication(w.AuthMidl.Authorization(w.UserGate.SignIn, adminOnly))).Methods("POST")

	router.HandleFunc("/api/nations", w.AuthMidl.Authentication(w.NationGate.GETAllNations)).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("spotter.davidebaldelli.it", "home.davidebaldelli.it"),
		Cache:      autocert.DirCache("certs"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Printf("Serving :7151 for domains: spotter.davidebaldelli.it, home.davidebaldelli.it")
	/*
		server2 := &http.Server{
			Addr:    ":7151",
			Handler: handler,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
	*/
	var wg sync.WaitGroup

	wg.Add(3)

	log.Printf("Serving :6316 for domains: home.davidebaldelli.it , api.acmodrepository.com")

	go func() {
		defer wg.Done()
		log.Fatal(server.ListenAndServeTLS("", ""))
	}()

	go func() {
		defer wg.Done()
		// serve HTTP, which will redirect automatically to HTTPS
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()
	/*
		go func() {
			defer wg.Done()
			log.Fatal(server2.ListenAndServeTLS("", ""))
		}()
	*/
	wg.Wait()

	//log.Fatal(http.ListenAndServe(":7151", handler))
}
