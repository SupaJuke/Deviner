package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKey string

func JWTKeyFunc(token *jwt.Token) (interface{}, error) {
	// "SupaJuke"
	if JWTKey == "" {
		log.Fatalln(errors.New("JWT key not found"))
	}

	return []byte(JWTKey), nil
}

func GetTokenFromHeader(r *http.Request) string {
	if _, tokenStr, ok := strings.Cut(r.Header.Get("Authentication"), "token "); ok {
		return tokenStr
	}

	return ""
}
