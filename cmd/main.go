package main

import (
	"log"
	"myapi/pkg/api"
	"myapi/pkg/repository"
	"myapi/pkg/tools"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	loadEnv()

	config := tools.New()
	psql, err := repository.New(config.Postgres.ConnStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	api := api.New("localhost:3004", mux.NewRouter(), psql)
	api.FillEndpoints()

	log.Fatal(api.ListenAndServe())
}
