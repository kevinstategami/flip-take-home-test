package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"flip-bank-statement-viewer/internal/model"
	"flip-bank-statement-viewer/internal/repository"
	"flip-bank-statement-viewer/internal/service"
)

// helper buat inisialisasi handler dengan repo/service kosong
func newTestHandler() *TransactionHandler {
	storage := []model.Transaction{}
	repo := repository.NewTransactionRepository(&storage)
	svc := service.NewTransactionService(repo)
	return NewTransactionHandler(svc)
}

func TestGetBalanceHandler(t *testing.T) {
	h := newTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/balance", nil)
	w := httptest.NewRecorder()

	h.GetBalance(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}

	// response harus JSON {"balance":0}
	var resp struct {
		Balance int64 `json:"balance"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response: %v; body: %s", err, w.Body.String())
	}

	if resp.Balance != 0 {
		t.Fatalf("expected balance 0 got %d", resp.Balance)
	}
}

func TestUpload_InvalidExtension(t *testing.T) {
	h := newTestHandler()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// buat file form dengan nama test.txt (invalid)
	fw, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		t.Fatalf("create form file failed: %v", err)
	}
	_, _ = fw.Write([]byte("dummy"))

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	h.Upload(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 invalid file type, got %d, body: %s", w.Code, w.Body.String())
	}

	// body harus JSON {"message": "..."}
	var errResp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &errResp); err != nil {
		t.Fatalf("expected json error message, got parsing err: %v", err)
	}

	msg, ok := errResp["message"]
	if !ok {
		t.Fatalf("error response missing message field: %s", w.Body.String())
	}
	if !strings.Contains(msg, "invalid file type") {
		t.Fatalf("unexpected error message: %s", msg)
	}
}

func TestUpload_Success(t *testing.T) {
	h := newTestHandler()

	// contoh CSV sesuai spec:
	csv := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
1624512883,COMPANY A,CREDIT,12000000,SUCCESS,salary
1624615065,E-COMMERCE B,DEBIT,150000,PENDING,clothes
`

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("file", "sample.csv")
	if err != nil {
		t.Fatalf("create form file failed: %v", err)
	}
	_, _ = io.Copy(fw, strings.NewReader(csv))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	h.Upload(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 created got %d, body: %s", w.Code, w.Body.String())
	}

	var successResp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &successResp); err != nil {
		t.Fatalf("failed to parse success response: %v ; body: %s", err, w.Body.String())
	}

	if msg, ok := successResp["message"]; !ok || msg != "Successfully uploaded transactions" {
		t.Fatalf("unexpected success message: %v", successResp)
	}

	// setelah upload, cek balance via handler
	bReq := httptest.NewRequest(http.MethodGet, "/balance", nil)
	bW := httptest.NewRecorder()
	h.GetBalance(bW, bReq)

	if bW.Code != http.StatusOK {
		t.Fatalf("expected balance 200 got %d", bW.Code)
	}

	var bResp struct {
		Balance int64 `json:"balance"`
	}
	if err := json.Unmarshal(bW.Body.Bytes(), &bResp); err != nil {
		t.Fatalf("failed to parse balance response: %v", err)
	}

	// expected balance: CREDIT (12000000) - DEBIT successful (250000) - other DEBIT successful?
	// From CSV: SUCCESS credits: 12000000 ; SUCCESS debits: 250000 => 11750000
	want := int64(11750000)
	if bResp.Balance != want {
		t.Fatalf("expected balance %d got %d", want, bResp.Balance)
	}
}

func TestGetIssuesHandler(t *testing.T) {
	h := newTestHandler()

	// first upload sample CSV like previous test
	csv := `1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
1624512883,COMPANY A,CREDIT,12000000,SUCCESS,salary
1624615065,E-COMMERCE B,DEBIT,150000,PENDING,clothes
`
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "sample.csv")
	if err != nil {
		t.Fatalf("create form file failed: %v", err)
	}
	_, _ = io.Copy(fw, strings.NewReader(csv))
	writer.Close()

	uploadReq := httptest.NewRequest(http.MethodPost, "/upload", body)
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())

	uploadW := httptest.NewRecorder()
	h.Upload(uploadW, uploadReq)

	if uploadW.Code != http.StatusCreated {
		t.Fatalf("upload failed in setup: status %d body %s", uploadW.Code, uploadW.Body.String())
	}

	// now request issues
	issuesReq := httptest.NewRequest(http.MethodGet, "/issues", nil)
	issuesW := httptest.NewRecorder()
	h.GetIssues(issuesW, issuesReq)

	if issuesW.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", issuesW.Code)
	}

	var issues []model.Transaction
	if err := json.Unmarshal(issuesW.Body.Bytes(), &issues); err != nil {
		t.Fatalf("failed to parse issues response: %v ; body: %s", err, issuesW.Body.String())
	}

	// from sample CSV there are 2 issues: one FAILED and one PENDING
	if len(issues) != 2 {
		t.Fatalf("expected 2 issues got %d", len(issues))
	}
}
