package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/handlers/gateways"
	"motorsportspotter.backend/handlers/middlewares"
	repo "motorsportspotter.backend/repositories/mysql"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

type Secret struct {
	Secret string
}

func main() {

	var cred Credentials

	if credFile, err := os.ReadFile("credentials.json"); err != nil {
		log.Fatal("no credentials file")
	} else {
		if err := json.Unmarshal(credFile, &cred); err != nil {
			log.Fatal("err parsing json")
		}
	}

	var secret Secret

	if secretFile, err := os.ReadFile("secret.json"); err != nil {
		log.Fatal("no secret file")
	} else {
		if err := json.Unmarshal(secretFile, &secret); err != nil {
			log.Fatal("err pasrsing json")
		}
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&loc=Local", cred.Username, cred.Password, cred.Host, "motorsport_spotter")
	dbase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database successfully")
	}

	router := handlers.Router{
		ChampGate:    gateways.ChampionshipGatewayImpl{Repo: repo.ChampionshipRepoImpl{Db: dbase}},
		TracksGate:   gateways.TracksGatewayImpl{Repo: repo.TracksRepositoryImpl{Db: dbase}},
		EventGate:    gateways.EventGatewayImpl{Repo: repo.EventsRepositoryImpl{Db: dbase}},
		SessionsGate: gateways.SessionsGatewayImpl{Repo: repo.SessionRepositoryImpl{Db: dbase}},
		NationGate:   gateways.NationsGatewayImpl{Repo: repo.NationsRepositoryImpl{Db: dbase}},
		UserGate:     gateways.UsersGatewayImpl{Repo: repo.UserRepositoryImpl{Db: dbase}, Secret: secret.Secret},
		AuthMidl:     middlewares.AuthorizationHandlerImpl{Secret: secret.Secret},
	}

	router.Listen()
}
