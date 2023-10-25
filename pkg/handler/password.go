package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"passwordManager/pkg/password"
)

func CreatePassword(w http.ResponseWriter, r *http.Request) {
	p := &password.Password{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		respond(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = password.CreatePassword(p)
	if err != nil {
		respond(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(w, "Password successfully created", http.StatusOK)
}

func GetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	p, err := password.GetPassword(vars["appname"])
	if err != nil {
		respond(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(w, p, http.StatusOK)
}
