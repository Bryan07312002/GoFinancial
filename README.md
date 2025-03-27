# Finance Organizer - Toy Project

A simple web application to organize personal finances, built as a learning project.

** frontend not fully working yet **

## Features

- Track income and expenses
- Categorize transactions
- Basic dashboard with financial overview
- SQLite database with GORM for ORM
- Backend API with Gorilla mux
- Vue.js frontend (prototype stage)

## Technologies Used

- **Backend**: Go (Golang)
  - Gorilla/mux for routing
  - GORM for ORM
  - SQLite database
- **Frontend**: Vue.js (prototype)
- **Build**: Go modules, npm/yarn for frontend

## Getting Started

### Prerequisites

- Go 1.23+
- Node.js (for frontend development)
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Bryan07312002/GoFinancial.git
cd GoFinancial
```

2. Set up sqlite3:
```bash
touch db.db
```

3. Set up backend:
```bash
go mod download
```

4. Set up frontend:
```bash
cd web
npm install
```

## Running the App

1. In one terminal
```bash
go run cmd/server/main.go
```

2. In another terminal
```bash
cd web
npm run dev
```
