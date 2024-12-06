package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type MyData struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

// For now, we use a hardcoded token.
var validToken = "my-secret-token"

func main() {
	http.Handle("/data", authMiddleware(http.HandlerFunc(dataHandler)))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := MyData{
		Message: "Hello from the server!",
		Value:   42,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing or invalid token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != validToken {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed
		next.ServeHTTP(w, r)
	})
}
