package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/SupaJuke/Indovinare/go/internal/models/response"
	"github.com/SupaJuke/Indovinare/go/internal/pkg/auth"
	"github.com/SupaJuke/Indovinare/go/internal/utils"
)

func Method(methods ...string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			// Handling pre-flight requests
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
				w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
				w.WriteHeader(http.StatusOK)
				return
			}

			supported := false
			for _, method := range methods {
				supported = r.Method == method
				if supported {
					break
				}
			}

			if !supported {
				log.Println("Unsupported method: ", r.Method)
				res := response.Response{
					Success: false,
					Msg:     "Method not supported. Expected " + strings.Join(methods, " "),
				}
				_ = res.WriteResp(w, http.StatusMethodNotAllowed)
				return
			}

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(hfn)
	}
}

func Authorize(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := utils.GetTokenFromHeader(r)
		httpCode, err := auth.Authorize(tokenStr)
		if err != nil {
			log.Println("User unauthorized")
			res := response.Response{
				Success: false,
				Msg:     fmt.Sprintf("Failed to login: %s", err),
			}
			_ = res.WriteResp(w, httpCode)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
