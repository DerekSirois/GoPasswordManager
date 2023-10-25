package handler

import (
	"encoding/json"
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
