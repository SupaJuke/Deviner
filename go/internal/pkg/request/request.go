package request

import (
	"log"
	"net/http"
	"strings"

	"github.com/SupaJuke/Indovinare/go/internal/pkg/auth"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func parseStatus(status int) string {
	switch status {
	case http.StatusUnauthorized:
		return "401 Unauthorized"
	case http.StatusBadRequest:
		return "400 Bad Request"
	default:
		return "500 Internal Server Error"
	}
}

func GetTokenFromHeader(r *http.Request) string {
	if _, tokenStr, ok := strings.Cut(r.Header.Get("Authentication"), "token "); ok {
		return tokenStr
	}

	return ""
}

func GetUsernameFromJWT(tokenStr string) string {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, &claims, auth.JWTKeyFunc)
	if err != nil {
		log.Println("Error while parsing claim: ", err)
		return parseStatus(http.StatusBadRequest)
	}
	return claims.Username
}
