package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/SupaJuke/Deviner/go/internal/handlers"
	"github.com/SupaJuke/Deviner/go/internal/middleware/auth"
)

// ================================ Middleware ================================ //

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
				w.WriteHeader(http.StatusMethodNotAllowed)
				_, err := w.Write([]byte("Method not supported. Expected " + strings.Join(methods, " ")))
				if err != nil {
					log.Panicln("Error while writing response [ValidateGet]", err)
				}
				return
			}

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(hfn)
	}
}

func Auth(handler http.Handler) http.Handler {
	return auth.Authorize(handler)
}

// ================================== Routes ================================== //

func Serve() {
	mux := http.NewServeMux()

	// Handlers
	loginHandler := http.HandlerFunc(handlers.Authenticate)
	guessHandler := http.HandlerFunc(handlers.CheckGuess)

	// Pre-authorized
	mux.Handle("/login", Method("POST")(loginHandler))

	// Post-authorized
	mux.Handle("/guess", Method("POST")(Auth(guessHandler)))

	// Serve
	log.Println("Now listening and serving on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Error while serving")
	}
}

// ============================================================================ //
