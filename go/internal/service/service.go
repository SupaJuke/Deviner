package service

import (
	"encoding/json"
	"log"
	"net/http"
	_ "os"
	"strings"
	"time"

	"github.com/SupaJuke/pooe-guessing-game/go/internal/accounts"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtKey = os.Getenv("JWT_KEY")
// This is from "SupaJuke"
var jwtKey = "1BA76F929C1AB69D6E0BA0AEC9F477ACC3563A9F65571B095E1D510AA20E9F62"

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

func Serve() {
	mux := http.NewServeMux()

	// Login
	mux.HandleFunc("/login", login)

	// Guess
	mux.Handle("/guess", authenticate(http.HandlerFunc(guess)))

	// Serve
	log.Println("Now listening and serving on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Error while serving")
	}
}

func getJWTKey(token *jwt.Token) (interface{}, error) {
	return []byte(jwtKey), nil
}

func getTokenFromHeader(r *http.Request) string {
	if _, tokenStr, ok := strings.Cut(r.Header.Get("Authentication"), "token "); ok {
		return tokenStr
	}
	return ""
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := getTokenFromHeader(r)
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, &claims, getJWTKey)
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

		// TODO: authorize here (check user against permission)
		// TODO2: maybe check curr username against username in token

		next.ServeHTTP(w, r)
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check method
	if method := r.Method; method != "POST" {
		log.Println("Method not supported: ", method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cred := Credentials{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cred); err != nil {
		log.Println("Error while parsing request: ", err)
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
	user, err := accounts.GetByUsername(cred.Username)
	if err != nil {
		log.Println("Failed while getting user: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	auth := user.Authenticate(cred.Password)
	if auth == nil {
		log.Println("Password Incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Genearting token for user
	expTime := time.Now().Add(30 * time.Minute)
	claims := Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Fatal("Internal error while generating token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(Response{Token: tokenStr})
	if err != nil {
		log.Fatal("Internal error while encoding response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

func guess(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("uwu"))
	if err != nil {
		return
	}
}
