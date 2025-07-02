package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func corsHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "service profile is up and running in docker!"}`)
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"username": "johndoe", "email": "john@example.com"}`)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7792"
	}
	
	http.HandleFunc("/", corsHandler(healthHandler))
	http.HandleFunc("/user/profile", corsHandler(userProfileHandler))
	
	fmt.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}