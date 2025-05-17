# Shift Scheduler App

A web application for managing employee shifts with separate interfaces for workers and admin. Built with Go (Gin) for the backend and Svelte for the frontend.

---

## 🔧 Technologies

- **Backend**: Go (Gin) + SQLite
- **Frontend**: Svelte
- **Docs**: Swagger UI
- **Containerization**: Docker & Docker Compose

---

## 🚀 Quick Start

### Prerequisites

- Go 1.23+
- Docker & Docker Compose
- Node.js (for frontend development)

### Run with Docker

```bash
docker-compose up --build
```

- Backend will be available at: [http://localhost:8080](http://localhost:8080)
- Swagger API Docs: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Frontend Login

Frontend will be available at: [http://localhost:3000](http://localhost:3000)

| Role   | Username | Password |
| ------ | -------- | -------- |
| Admin  | admin    | admin    |
| Worker | justin   | user123  |
| Worker | payd     | user123  |
