package main

import (
	"log"
	"net/http"
	"os"

	"golang-proto/producer/route"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("general.env")
	if err != nil {
		log.Fatal("Error loading general.env file")
	}

	log.Printf("Starting GoLang Server")
	router := route.NewRouter()
	log.Printf("Router created")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
