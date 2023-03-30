package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SupaJuke/Indovinare/go/internal/models/response"
	"github.com/SupaJuke/Indovinare/go/internal/pkg/auth"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	// Parsing request
	cred := Credentials{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cred); err != nil {
		log.Println("Error while parsing request [Authenticate]: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr, httpCode, err := auth.Authenticate(cred.Username, cred.Password)
	if err != nil {
		res := response.Response{
			Success: false,
			Msg:     fmt.Sprintf("Failed to login: %s", err),
		}
		err = res.WriteResp(w, httpCode)
		if err == nil {
			log.Printf("User '%s' failed to login", cred.Username)
		}
		return
	}

	res := response.Response{
		Success: true,
		Msg:     "Successfully logged in",
		Token:   tokenStr,
	}
	err = res.WriteResp(w, httpCode)
	if err == nil {
		log.Printf("User '%s' logged in successfully", cred.Username)
	}
}
