package handlers

import (
	"net/http"

	"github.com/SupaJuke/Indovinare/go/internal/middleware/auth"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	auth.Authenticate(w, r)
}
