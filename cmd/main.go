package main

import (
	"log"
	"myapi/pkg/api"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// api := api.New("localhost:3001", http.NewServeMux())
	// api.FillEndpoints()
	// log.Fatal(api.ListenAndServe())
	loadEnv()

	api.TestDB()
}
