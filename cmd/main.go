package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", HelloHandler)

	log.Fatal(http.ListenAndServe("localhost:3001", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := w.Write([]byte("hello"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		_, err := w.Write([]byte("hello from POST"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		_, err := w.Write([]byte("hello from another method"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
