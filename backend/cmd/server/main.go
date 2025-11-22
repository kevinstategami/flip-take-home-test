package main

import (
	"log"
	"net/http"

	"flip-bank-statement-viewer/internal/handler"
	"flip-bank-statement-viewer/internal/repository"
	"flip-bank-statement-viewer/internal/service"
	"flip-bank-statement-viewer/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	repo := repository.NewTransactionRepository(&store.Transactions)
	svc := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("/upload", h.Upload)
	mux.HandleFunc("/balance", h.GetBalance)
	mux.HandleFunc("/issues", h.GetIssues)

	handlerWithCORS := enableCORS(mux)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", handlerWithCORS)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
