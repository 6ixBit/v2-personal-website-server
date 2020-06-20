package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/6ixBit/v2-Personal-Website/handlers"
	mwr "github.com/6ixBit/v2-Personal-Website/middleware"
	"github.com/joho/godotenv"
)

func init() {
	loadEnvFile()
}

func tt() {
	fmt.Println("Gimme mioney rn")
	fmt.Println("we out here")
}

func main() {
	http.HandleFunc("/", 		 mwr.LogRequests(handlers.HomeHandler))
	http.HandleFunc("/contact",  mwr.LogRequests(handlers.ContactHandler))
	http.HandleFunc("/projects", mwr.LogRequests(handlers.ProjectsHandler))

	startServer()
}

func startServer() {
	fmt.Println("Server up and running on port 8080..")
	http.ListenAndServe(":8080", nil)
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println("Env File Error: Failed to load .env file with environment variable")
	}
}

