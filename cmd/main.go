package main

import (
	"log"
	"myapi/pkg/api"
	"myapi/pkg/tools"
	"net/http"

	"strings"

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
	serverUrl := strings.Join([]string{config.Server.Host, ":", config.Server.Port}, "")
	api := api.New(serverUrl, http.NewServeMux())
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe(serverUrl))
	// api.TestDB()
}
