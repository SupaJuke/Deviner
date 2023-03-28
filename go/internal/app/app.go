package app

import (
	"log"
	"os"

	db "github.com/SupaJuke/Indovinare/go/internal/database"
	"github.com/SupaJuke/Indovinare/go/internal/middleware/auth"
	"github.com/SupaJuke/Indovinare/go/internal/router"

	"github.com/joho/godotenv"
)

// const defaultPort = "8080"

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db.InitDB()
	log.Println("Initialized connection to database")
	auth.JWTKey = os.Getenv("JWT_KEY")
	router.Serve()
}
