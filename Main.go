package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	http2 "motorsportspotter.backend/controllers/http"
	repo "motorsportspotter.backend/repositories/mysql"
)

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&loc=Local", "root", "SP589a%6", "api.acmodrepository.com", "motorsport_spotter")
	dbase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database successfully")
	}

	router := http2.Web{
		ChampCtrl:    http2.ChampionshipControllerImpl{Repo: repo.ChampionshipRepoImpl{Db: dbase}},
		TracksCtrl:   http2.TracksControllerImpl{Repo: repo.TracksRepositoryImpl{Db: dbase}},
		EventCtrl:    http2.EventControllerImpl{Repo: repo.EventsRepositoryImpl{Db: dbase}},
		SessionsCtrl: http2.SessionsControllerImpl{Repo: repo.SessionRepositoryImpl{Db: dbase}},
		NewsCtrl:     http2.NewsControllerImpl{Repo: repo.NewsRepositoryImpl{Db: dbase}},
	}

	router.Listen()
}
