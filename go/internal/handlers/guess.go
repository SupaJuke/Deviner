package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Guess struct {
	Guess string `json:"guess"`
}

func CheckGuess(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("uwu"))
	if err != nil {
		return
	}

	// TODO: implement this
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	guess := Guess{}
	if err := dec.Decode(&guess); err != nil {
		log.Println("Error while parsing request: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
