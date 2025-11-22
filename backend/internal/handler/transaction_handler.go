package handler

import (
	"encoding/json"
	"flip-bank-statement-viewer/internal/service"
	"flip-bank-statement-viewer/internal/utils"
	"net/http"
	"path/filepath"
	"strings"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func NewTransactionHandler(svc service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: svc}
}

func (h *TransactionHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		writeError(w, http.StatusBadRequest, "failed to parse form: "+err.Error())
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid file upload")
		return
	}
	defer file.Close()

	filename := header.Filename
	ext := strings.ToLower(filepath.Ext(filename))

	if ext != ".csv" {
		writeError(w, http.StatusBadRequest, "invalid file type: only .csv allowed")
		return
	}

	data, err := utils.ParseCSV(file)
	if err != nil {
		writeError(w, http.StatusBadRequest, "failed to parse csv: "+err.Error())
		return
	}

	if err := h.svc.Upload(data); err != nil {
		writeError(w, http.StatusBadRequest, "validation error: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully uploaded transactions",
	})
}

func (h *TransactionHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	balance := h.svc.GetBalance()

	response := struct {
		Balance int64 `json:"balance"`
	}{
		Balance: balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) GetIssues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	issues := h.svc.GetIssues()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"message": msg,
	})
}
