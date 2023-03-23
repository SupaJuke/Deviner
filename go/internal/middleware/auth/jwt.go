package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	// "SupaJuke"
	JWTKey := os.Getenv("JWT_KEY")
	if JWTKey == "" {
		log.Fatal(errors.New("JWT key not found"))
	}

	return []byte(JWTKey), nil
}

func getKey() []byte {
	return []byte(os.Getenv("JWT_KEY"))
}

func getTokenFromHeader(r *http.Request) string {
	if _, tokenStr, ok := strings.Cut(r.Header.Get("Authentication"), "token "); ok {
		return tokenStr
	}

	return ""
}
