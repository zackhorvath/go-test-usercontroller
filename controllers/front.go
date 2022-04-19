package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)  // requests to be handled on the /users route, and pass in *uc object
	http.Handle("/users/", *uc) // handlers are explicit in their pathing
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
