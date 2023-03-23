package handlers

import (
	"net/http"

	"github.com/SupaJuke/Deviner/go/internal/middleware/auth"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	auth.Authenticate(w, r)
}
