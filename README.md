# Warlock Backend

A backend service written in Go for managing the server side tasks for Warlock AI. My main purpose with this project is to get familiar with Go and it's ecosystem.

## 📦 Dependencies

See [`go.mod`](./go.mod) for full list. Notable ones:

- `github.com/gin-gonic/gin`
- `gorm.io/gorm` and `gorm.io/driver/postgres`

## 🛠️ Installation

```bash
git clone https://github.com/tamlaz/warlock-backend.git
cd warlock-backend
go mod download
go run main.go

