package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"sync"
)

type URLStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

var store = &URLStore{
	urls: make(map[string]string),
}

type ShortRequest struct {
	URL string `json:"url"`
}

type ShortResponse struct {
	ShortURL string `json:"short_url"`
}

func generateId() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)[:6]
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	var req ShortRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	id := generateId()

	store.mu.Lock()
	store.urls[id] = req.URL
	store.mu.Unlock()

	shortURL := fmt.Sprintf("http://localhost:8080/%s", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortResponse{ShortURL: shortURL})

}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if id == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	store.mu.RLock()
	longURL, exists := store.urls[id]
	store.mu.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
			appEnv = "development" 
	}

	fmt.Printf("Starting server in [%s] mode ...\n", appEnv)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealth)

	mux.HandleFunc("POST /shorten", handleShorten)

	mux.HandleFunc("GET /", handleRedirect)


	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
