# Warlock Backend

A backend service written in Go for managing the server-side tasks for Warlock AI.  
My main purpose with this project is to get familiar with Go and its ecosystem.

## 📦 Dependencies

See [`go.mod`](./go.mod) for the full list. Notable ones:

- `github.com/gin-gonic/gin`
- `gorm.io/gorm` and `gorm.io/driver/postgres`

## 🛠️ Installation

```bash
git clone https://github.com/tamlaz/warlock-backend.git
cd warlock-backend
go mod download
go run main.go
```

---

## 📡 API Endpoints

All routes below assume the server is running locally on **`http://localhost:8080`**.

| Method | Path                         | Purpose                            | Auth required
|--------|------------------------------|------------------------------------|--------------
| GET    | `/health`                          | Health-check (“Is backend alive?”) | ❌           
| POST   | `/api/go/v1/signup`          | Register a new user                | ❌           
| POST   | `/api/go/v1/login`           | Authenticate user & receive JWT    | ❌           
| POST   | `/api/go/v1/validate-user`   | Validate a JWT and fetch user info | ✅ (JWT)     

---

### 1. `GET /`

Health-check endpoint.

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "message": "Warlock backend is up and running"
}
```

---

### 2. `POST /api/go/v1/signup`

Creates a new account.

```bash
curl -X POST http://localhost:8080/api/go/v1/signup \
  -H "Content-Type: application/json" \
  -d '{
        "email": "alice@example.com",
        "password": "Str0ngP@ssw0rd!",
        "firstName": "Alice",
        "lastName": "Wonder"
      }'
```

**Response:**
```json
{
  "message": "User created successfully"
}
```

---

### 3. `POST /api/go/v1/login`

Authenticates a user and returns a signed JWT.

```bash
curl -X POST http://localhost:8080/api/go/v1/login \
  -H "Content-Type: application/json" \
  -d '{
        "email": "alice@example.com",
        "password": "Str0ngP@ssw0rd!"
      }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

---

### 4. `POST /api/go/v1/validate-user`

Validates a previously-issued JWT and returns user details.

```bash
curl -X POST http://localhost:8080/api/go/v1/validate-user \
  -H "Content-Type: application/json" \
  -d '{
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
      }'
```

**Response:**
```json
{
  "userId": 1,
  "roles": [
    {
      "name": "Student"
    }
  ]
}
```

---

### 🔐 Working with JWTs locally

1. Sign up ➜ copy the email/password you used.  
2. **Login** to get a `token`.  
3. Supply that token in the JSON body for `validate-user` (or in an `Authorization: Bearer <token>` header after modifying the controller).
