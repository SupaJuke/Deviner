package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/SupaJuke/Indovinare/go/internal/models/response"
	"github.com/SupaJuke/Indovinare/go/internal/models/users"
	"github.com/SupaJuke/Indovinare/go/internal/utils"
)

type Guess struct {
	Guess string `json:"guess"`
}

func compareGuess(username, guess string) (int, error) {
	user, err := users.GetByUsername(username)
	if err != nil {
		log.Println("Cannot get user: ", err)
		return http.StatusInternalServerError, err
	}

	// Fetch code from DB
	code, err := user.GetCode()
	if err != nil {
		log.Println("Failed to get user's code: ", err)
		return http.StatusInternalServerError, err
	}

	// Validating answer
	p := "^[0-9]{5}$"
	matched, err := regexp.MatchString(p, guess)
	if err != nil {
		log.Println("Error while matching regexp pattern")
		return http.StatusBadRequest, err
	}
	if !matched {
		msg := "Request must be from 00000 to 99999"
		log.Println(msg + ": " + guess)
		return http.StatusBadRequest, errors.New(msg)
	}

	// Check if guess correct
	if guess != code {
		yDigit, gDigit := 0, 0
		for i := range code {
			if code[i] == guess[i] {
				gDigit++
				code = strings.Replace(code, string(code[i]), "x", 1)
			}
		}
		for i := range code {
			if strings.Contains(code, string(guess[i])) {
				yDigit++
				code = strings.Replace(code, string(guess[i]), "x", 1)
			}
		}
		errMsg := fmt.Sprintf("G: %d Y: %d", gDigit, yDigit)
		return http.StatusOK, errors.New(errMsg)
	}

	// Generate new code
	err = user.GenerateNewCode()
	if err != nil {
		log.Println("Failed to generate code: ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func HandleGuess(w http.ResponseWriter, r *http.Request) {
	// Parsing request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	guess := Guess{}
	if err := dec.Decode(&guess); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing request [HandleGuess]: ", err)
		return
	}

	// Extracting guess
	username := utils.GetUsernameFromJWT(utils.GetTokenFromHeader(r))
	httpCode, err := compareGuess(username, guess.Guess)

	// Creating response
	if err != nil {
		res := response.Response{
			Success: false,
			Msg:     fmt.Sprintf("Guess failed: %s", err),
		}
		if httpCode == http.StatusOK {
			var g, y string
			fmt.Sscanf(err.Error(), "G: %s Y: %s", &g, &y)
			res.Msg = "Guess incorrect"
			res.Green = g
			res.Yellow = y
		}

		err = res.WriteResp(w, httpCode)
		if err == nil {
			log.Printf("Guess unsuccessful for user '%s'", username)
		}
		return
	}

	// Create res
	res := response.Response{
		Success: true, Msg: "Guess is correct; regenerating new code",
	}
	err = res.WriteResp(w, httpCode)
	if err == nil {
		log.Printf("User '%s' guessed correctly", username)
	}
}
