package controllers

import (
	"encoding/json"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(data)
}
