package api

import (
	"encoding/json"
	"myapi/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *api) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		stringID, ok := vars["id"]
		if ok {
			id, err := strconv.Atoi(stringID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			data, err := api.psql.GetUserByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			err = json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			return
		}

		data, err := api.psql.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	case http.MethodPost:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = api.psql.NewUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(http.StatusCreated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	case http.MethodDelete:
		vars := mux.Vars(r)
		stringID, ok := vars["id"]
		if ok {
			id, err := strconv.Atoi(stringID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			err = api.psql.DeleteUser(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			err = json.NewEncoder(w).Encode(http.StatusOK)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			return
		}

		json.NewEncoder(w).Encode(http.StatusBadRequest)
	}
}
