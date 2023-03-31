package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/SupaJuke/Indovinare/go/internal/models/users"
	"github.com/SupaJuke/Indovinare/go/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Authenticate(username, password string) (string, int, error) {
	// Check username not empty
	if username == "" {
		log.Println("Username empty: ", username)
		return "", http.StatusBadRequest, errors.New("Username empty")
	}

	// Check username exists in db
	user, err := users.GetByUsername(username)
	if err != nil {
		log.Println("Failed while getting user: ", err)
		return "", http.StatusOK, errors.New("User not found")
	}

	// Check password
	auth := user.Authenticate(password)
	if auth != nil {
		log.Printf("User %s Password Incorrect", username)
		return "", http.StatusOK, errors.New("Username or Password incorrect")
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
	tokenStr, err := token.SignedString([]byte(utils.JWTKey))
	if err != nil {
		log.Println("Internal error while generating token: ", err)
		return "", http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return tokenStr, http.StatusOK, nil

	// NOTE: implement cookie down the line?
	/*
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}
	*/
}

func Authorize(tokenStr string) (int, error) {
	if tokenStr == "" {
		return http.StatusBadRequest, errors.New("Token not found")
	}

	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, utils.JWTKeyFunc)
	if err != nil {
		log.Println("Failed after parsing claims:", err)
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, errors.New("Failed reading token")
	}

	if !token.Valid {
		return http.StatusUnauthorized, errors.New("Token invalid")
	}

	return http.StatusOK, nil

	// TODO: maybe check curr username against username in token
	// TODO: check user permission
}
