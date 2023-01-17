package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"motorsportspotter.backend/controllers/http"
	repo "motorsportspotter.backend/repositories/mysql"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func main() {

	var cred Credentials

	if jsonFile, err := os.ReadFile("credentials.json"); err != nil {
		log.Fatal("no credentials file")
	} else {
		if err := json.Unmarshal(jsonFile, &cred); err != nil {
			log.Fatal("err parsing json")
		}
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&loc=Local", cred.Username, cred.Password, cred.Host, "motorsport_spotter")
	dbase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database successfully")
	}

	router := http.Web{
		ChampCtrl:    http.ChampionshipControllerImpl{Repo: repo.ChampionshipRepoImpl{Db: dbase}},
		TracksCtrl:   http.TracksControllerImpl{Repo: repo.TracksRepositoryImpl{Db: dbase}},
		EventCtrl:    http.EventControllerImpl{Repo: repo.EventsRepositoryImpl{Db: dbase}},
		SessionsCtrl: http.SessionsControllerImpl{Repo: repo.SessionRepositoryImpl{Db: dbase}},
		NewsCtrl:     http.NewsControllerImpl{Repo: repo.NewsRepositoryImpl{Db: dbase}},
	}

	router.Listen()
}
