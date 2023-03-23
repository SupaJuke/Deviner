package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SupaJuke/Deviner/go/internal/middleware/auth"
	"github.com/SupaJuke/Deviner/go/internal/models/users"
	"github.com/SupaJuke/Deviner/go/internal/utils"
)

type Guess struct {
	Guess string `json:"guess"`
}

type Response struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"message,omitempty"`
}

func CheckGuess(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	guess := Guess{}
	if err := dec.Decode(&guess); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing request [CheckGuess]: ", err)
		return
	}

	tokenStr := auth.GetTokenFromHeader(r)
	username := utils.GetUsernameFromJWT(tokenStr)
	log.Println("username :", username)
	user, err := users.GetByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Cannot get user: ", err)
		return
	}

	code, err := user.GetCode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to get user's code: ", err)
		return
	}

	// Check if guess correct
	if guess.Guess == code {
		// Generate new code
		err = user.GenerateNewCode()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Failed to generate code: ", err)
			return
		}

		// Create res
		res, err := json.Marshal(Response{Success: true, Msg: "Guess is correct; regenerating new code"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Internal error while encoding response: ", err)
			return
		}

		// Write res
		_, err = w.Write(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Internal error while writing response: ", err)
			return
		}

		return
	}

	// Create res
	res, err := json.Marshal(Response{Success: false, Msg: "Guess is incorrect"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal error while encoding response: ", err)
		return
	}

	// Write res
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal error while writing response: ", err)
		return
	}
}
