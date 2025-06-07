package main

import (
	"log"
	"myapi/pkg/api"
	"net/http"
)

func main() {
	api := api.New("localhost:3001", http.NewServeMux())
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe())
}
