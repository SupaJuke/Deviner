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
