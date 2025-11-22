# Frontend (Next.js) - Flip Bank Statement Viewer

This is the **Next.js frontend** for the Flip Bank Statement Viewer project.  
It provides a clean UI for uploading CSV transactions, viewing the balance summary, and listing problematic transactions.

The frontend follows a **modular Clean Architecture**, separating UI components, hooks, services, and utilities.

---

# 🚀 Features

- Upload CSV to backend
- Auto-redirect after successful upload
- Display computed balance
- Display FAILED & PENDING transactions
- Native CSS (no Tailwind)
- Branded Flip color palette
- Uses formatters (currency, names, status, date)
- Server Component fetch optimization optional

---

# 📁 Folder Structure

```
frontend/
│
├── src/
│   ├── app/
│   │   ├── page.tsx                  → redirects to /upload
│   │   ├── upload/page.tsx           → upload flow routing
│   │   └── transactions/page.tsx     → summary + issues
│   │   └── globals.css
|   |   └── page.tsx
│   ├── modules/
│   │   ├── upload/
│   │   │   ├── components/           → UploadForm
│   │   │   └── hooks/                → useUpload()
│   │   └── transactions/
│   │       ├── components/           → BalanceSummary, TransactionsTable
│   │       └── hooks/                → useTransactions()
│   │
│   ├── services/
│   │   └── api.ts                    → API client
│   │
│   ├── utils/
│   │   └── formatter.ts              → capitalize, currency, status, date
│   │
│   ├── app/globals.css               → global CSS + brand colors
│   └── public/
│
└── package.json
```

---

# 🧠 Architecture Decisions

## 1. Pages Only Handle Routing  
Example:  
`/upload` page only calls `<UploadForm />`  
`/transactions` page only renders components  
→ No logic inside pages, clean and thin.

---

## 2. Components Handle UI Only  
Components do NOT fetch data.  
Example:

- `UploadForm` → UI only
- `BalanceSummary` → just display data
- `TransactionsTable` → just display rows

---

## 3. Hooks Handle Logic  
Hooks handle state, API calls, and validation.

- `useUpload()` → upload flow  
- `useTransactions()` → fetch balance + issues

Separation makes each layer testable & replaceable.

---

## 4. Services Handle API Requests  
All API requests are centralized in `src/services/api.ts`.

Also includes error normalization:

```json
{ "message": "invalid csv format" }
```

---

## 5. Utility Formatters  
Formatters are placed in `src/utils/formatter.ts`:

- capitalize
- capitalizeWords
- formatCurrency
- formatUnixDate
- formatStatus
- formatAmount

Keeps UI components clean.

---

## 6. Native CSS With Flip Branding  
Defined in `globals.css`:

- `--flip-primary`
- `--flip-primary-dark`
- `--flip-accent`
- `--error`
- `--warning`

UI is styled similar to Flip’s aesthetic.

---

# 🔧 Setup Instructions

## 1. Install Dependencies

```sh
cd frontend
npm install
```

---

## 2. Environment Variables

Create:

### `.env.local`
```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

Docker will override this with:

```
NEXT_PUBLIC_API_URL=http://backend:8080
```

---

## 3. Development Server

```sh
npm run dev
```

Frontend runs at:

👉 http://localhost:3000

---

## 4. Build for Production

```sh
npm run build
npm start
```

---

# 🐳 Docker Support

Frontend is fully containerized via `/frontend/Dockerfile`.

Used by root-level `docker-compose.yml` with:

```yaml
environment:
  NEXT_PUBLIC_API_URL: "http://backend:8080"
```

Then run everything:

```sh
docker compose up --build
```

---

# ⚙️ Scripts

```json
"scripts": {
  "dev": "next dev",
  "build": "next build",
  "start": "next start"
}
```

---

# 🧪 Testing (Optional)

React Testing Library + Vitest can be added easily (not required in assignment).

---

# 👨‍💻 Author  
Kelvin Febrian  
Frontend / Fullstack Engineer (Next.js + Go)
