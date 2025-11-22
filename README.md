# Flip Bank Statement Viewer

A full-stack application built using **Golang** (backend) and **Next.js** (frontend), designed to upload CSV bank statements, validate data, compute balances, and show transaction issues.

The project follows **clean architecture**, is fully containerized using **Docker Compose**, and includes CI automation via **GitHub Actions**.

---

# 🚀 Features

- CSV upload & validation (extension + format)
- Compute final balance (CREDIT–DEBIT, SUCCESS only)
- Detect problematic transactions (FAILED & PENDING)
- Consistent error response format (`{ "message": "error" }`)
- Clean Architecture (both Backend & Frontend)
- Dockerized: single command to run full stack
- GitHub Actions CI for test + build

---

# 📦 Folder Structure

```
flip-bank-statement-viewer/
│
├── backend/       → Golang service
├── frontend/      → Next.js client
└── docker-compose.yml
```

---

# 🛠️ Setup Instructions

## 1. Requirements
- Go 1.23+
- Node.js 20+
- Docker & Docker Compose (recommended)
- Git

---

## 2. Running Locally (No Docker)

### Backend
```sh
cd backend
go mod download
go run cmd/server/main.go
```
Runs at → **http://localhost:8080**

### Frontend
```
cd frontend
npm install
npm run dev
```
Runs at → **http://localhost:3000**

Make sure env:

`frontend/.env.local`:
```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

---

## 3. Running With Docker (Recommended)

At project root:

```sh
docker compose build
docker compose up
```

Services:
- Frontend → http://localhost:3000
- Backend → http://localhost:8080

To stop everything:
```sh
docker compose down
```

---

## 4. Running Backend Tests

```sh
cd backend
go test ./... -v
```

Includes:
- Handler tests
- Service tests
- Error format tests

---

# 🏗️ Architecture Decisions

This project is designed using **Clean Architecture** principles to maximize clarity, modularity, and separation of concerns.

---

## 🧩 Backend Architecture (Golang)

```
backend/
│
├── cmd/server/           → App entrypoint (HTTP server)
├── internal/
│   ├── handler/          → HTTP handlers (request/response)
│   ├── service/          → Business logic
│   ├── repository/       → In-memory data storage
│   ├── model/            → Domain models + enums
│   └── utils/            → CSV parser
```

### Key decisions

### **1. Separated Layers**
- **Handler**: only HTTP + JSON formatting  
- **Service**: validation + domain rules  
- **Repository**: independent storage provider  
- **Model**: strong typed domain with enums  

This makes unit testing simple and avoids mixing logic.

---

### **2. Consistent Error Format**

All errors returned as:

```json
{
  "message": "error description"
}
```

Implemented using helper:

```go
func writeError(w, status, msg)
```

This ensures frontend always receives predictable JSON.

---

### **3. CSV Validation**

Backend validates:

- File extension: must be `.csv`
- CSV row count: exactly 6 fields
- Allowed types: `DEBIT`, `CREDIT`
- Allowed status: `SUCCESS`, `FAILED`, `PENDING`
- Non-negative amount

Service layer returns error:
```
validation error: invalid transaction type at row 2
```

---

## 🎨 Frontend Architecture (Next.js)

```
frontend/
│
├── src/app/                       → Routing only
│   ├── upload/
│   ├── transactions/
│   └── page.tsx (redirect to /upload)
│
├── src/modules/
│   ├── upload/
│   │   ├── components/
│   │   └── hooks/
│   └── transactions/
│       ├── components/
│       └── hooks/
│
├── src/services/                  → API client
└── src/utils/                     → formatters, helpers
```

### Key decisions

### **1. Pages handle routing only**
Pages do **not** contain logic or UI.  
They simply render components from `modules/`.

### **2. Components are UI-only**
No business logic.  
No HTTP calls.

### **3. Hooks handle all state + fetching**
Example:
- `useUpload()` → upload logic
- `useTransactions()` → fetch balance + issues

### **4. Strict Separation: UI | Logic | Transport**
This mirrors backend clean architecture one-to-one.

### **5. First fetch always from server component (optional optimization)**
To avoid double-fetch in development, server components can preload data.

---

## 🐳 Docker Architecture

### **Separate Dockerfiles**
- `backend/Dockerfile` → Go build
- `frontend/Dockerfile` → Next.js multi-stage build

### **docker-compose.yml** runs both:

- Backend → exposes **8080**
- Frontend → exposes **3000**
- Internal network: frontend connects to backend via hostname `backend`

### Example:
```yaml
environment:
  NEXT_PUBLIC_API_URL: "http://backend:8080"
```

---

## 🤖 CI/CD (GitHub Actions)

Workflow includes:

- Backend: `go test`, `go build`
- Frontend: `npm install`, `npm run build`
- Docker Compose build validation

Workflow triggers only on changes inside:

```
backend/**
frontend/**
docker-compose.yml
```

This keeps pipeline efficient for monorepo.

---

# 📄 License
This project is provided as part of a take-home exercise.

---

# ✨ Author
Kelvin Febrian  
Fullstack Engineer (Go & Next.js)
