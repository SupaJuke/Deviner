package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SupaJuke/Deviner/go/internal/models/users"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token   string `json:"token,omitempty"`
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"message,omitempty"`
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cred := Credentials{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cred); err != nil {
		log.Println("Error while parsing request [Authenticate]: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check username not empty
	if cred.Username == "" {
		log.Println("Username empty: ", cred.Username)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check username exists in db
	user, err := users.GetByUsername(cred.Username)
	if err != nil {
		log.Println("Failed while getting user: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	auth := user.Authenticate(cred.Password)
	if auth != nil {
		log.Println("Password Incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Genearting token for user
	expTime := time.Now().Add(24 * time.Hour)
	claims := Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		log.Println("Internal error while generating token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(Response{Token: tokenStr})
	if err != nil {
		log.Println("Internal error while encoding response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal error while writing response: ", err)
		return
	}

	log.Printf("User '%s' logged in successfully", cred.Username)

	// NOTE: implement cookie down the line?
	/*
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}
	*/
}

func Authorize(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := GetTokenFromHeader(r)
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, &claims, JWTKeyFunc)
		if err != nil {
			log.Println("failed after parsing claims:", err)
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: maybe check curr username against username in token
		// TODO: check user permission

		handler.ServeHTTP(w, r)
	})
}
