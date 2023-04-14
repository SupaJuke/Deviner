package app

import (
	"log"
	"os"

	db "github.com/SupaJuke/Indovinare/go/internal/database"
	"github.com/SupaJuke/Indovinare/go/internal/pkg/auth"
	"github.com/SupaJuke/Indovinare/go/internal/router"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db.InitDB()
	log.Println("Initialized connection to database")
	// _, _ = users.Create("tester", "password")
	auth.JWTKey = os.Getenv("JWT_KEY")
	router.Serve()
}
