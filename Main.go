package main

import (
	"fmt"
	"github.com/joho/godotenv"
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
	DbName   string `json:"dbName"`
}

type Secret struct {
	Secret string
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cred Credentials

	if dbUsername := os.Getenv("DB_USERNAME"); dbUsername == "" {
		log.Fatal("missing DB_USERNAME in env")
	} else {
		cred.Username = dbUsername
	}

	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword == "" {
		log.Fatal("missing DB_PASSWORD in env")
	} else {
		cred.Password = dbPassword
	}

	if dbName := os.Getenv("DB_DATABASE"); dbName == "" {
		log.Fatal("missing DB_DATABASE in env")
	} else {
		cred.DbName = dbName
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost == "" {
		log.Fatal("missing DB_HOST in env")
	} else {
		cred.Host = dbHost
	}

	var secret Secret

	if secretKey := os.Getenv("SECRET_KEY"); secretKey == "" {
		log.Fatal("missing SECRET_KEY in env")
	} else {
		secret.Secret = secretKey
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&loc=Local", cred.Username, cred.Password, cred.Host, cred.DbName)
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
