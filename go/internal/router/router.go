package router

import (
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Now listening and serving on port :" + port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalln("Error while serving")
	}
}

// ============================================================================ //
