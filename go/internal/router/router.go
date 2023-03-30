package router

import (
	"log"
	"net/http"

	h "github.com/SupaJuke/Indovinare/go/internal/handlers"
	mw "github.com/SupaJuke/Indovinare/go/internal/middleware"
)

// ================================== Routes ================================== //

func Serve() {
	mux := http.NewServeMux()

	// Handlers
	loginHandler := http.HandlerFunc(h.HandleAuthenticate)
	guessHandler := http.HandlerFunc(h.HandleGuess)

	// Pre-authorized
	mux.Handle("/login", mw.Method("POST")(loginHandler))

	// Post-authorized
	mux.Handle("/guess", mw.Method("POST")(mw.Authorize(guessHandler)))

	// Serve
	log.Println("Now listening and serving on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Error while serving")
	}
}

// ============================================================================ //
