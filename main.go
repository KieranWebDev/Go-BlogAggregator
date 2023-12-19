package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment variables")
	}
	fmt.Println("PORT", portString)

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Printf("Starting server at port %s\n", portString)
	srv.ListenAndServe()

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
