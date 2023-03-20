package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SupaJuke/pooe-guessing-game/go/internal/accounts"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_KEY")

// NOTE: can use type "Credentials" instead of "User"
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
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
	log.Fatal("Error while serving: ", http.ListenAndServe(":8080", mux))
}

func getJWTKey(token *jwt.Token) (interface{}, error) {
	return jwtKey, nil
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
		token, err := jwt.ParseWithClaims(tokenStr, claims, getJWTKey)
		if err != nil {
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

		next.ServeHTTP(w, r)
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check method
	if method := r.Method; method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	pwd := r.FormValue("pwd")
	// Check username not empty
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check username exists in db
	user, err := accounts.GetByUsername(username)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	auth := user.Authenticate(pwd)
	if auth == nil {
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
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(Response{Token: tokenStr})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// NOTE: implement cookie down the line?
	/*
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}
	*/
}

func guess(w http.ResponseWriter, r *http.Request) {
}
