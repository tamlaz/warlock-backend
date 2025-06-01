# Warlock Backend

A backend service written in Go for managing the server-side tasks for Warlock AI.  
My main purpose with this project is to get familiar with Go and its ecosystem.

## 📦 Dependencies

See `go.mod` for the full list. Notable ones: github.com/gin-gonic/gin, gorm.io/gorm, gorm.io/driver/postgres, github.com/gorilla/websocket

## 🛠️ Installation

Clone the repo, download dependencies, and run:

git clone https://github.com/tamlaz/warlock-backend.git  
cd warlock-backend  
go mod download  
go run main.go

## 📡 API Endpoints

All routes assume the server is running locally on http://localhost:8080.

| Method | Path                                | Purpose                               | Auth required 
|--------|-------------------------------------|---------------------------------------|---------------
| GET    | /health                             | Health-check (“Is backend alive?”)    | ❌            
| POST   | /api/go/v1/signup                   | Register a new user                   | ❌            
| POST   | /api/go/v1/login                    | Authenticate user & receive JWT       | ❌            
| GET   | /api/go/v1/validate-user            | Validate a JWT and fetch user info    | ❌      
| PUT    | /api/go/v1/add-strike-to-user       | Increment a user's strike count       | ❌    
| GET    | /ws                                 | Subscribe to real-time ban events     | ❌            

### GET /health

Performs a simple health check.

curl http://localhost:8080/health

Response:  
{ "message": "Warlock backend is up and running" }

### POST /api/go/v1/signup

Registers a new user.

curl -X POST http://localhost:8080/api/go/v1/signup \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "Str0ngP@ssw0rd!", "firstName": "Alice", "lastName": "Wonder"}'

Response:  
{ "message": "User created successfully" }

### POST /api/go/v1/login

Authenticates a user and returns a signed JWT.

curl -X POST http://localhost:8080/api/go/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "Str0ngP@ssw0rd!"}'

Response:  
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..." }

### POST /api/go/v1/validate-user

Validates a JWT and fetches user information.

curl -X POST http://localhost:8080/api/go/v1/validate-user \
  -H "Content-Type: application/json" \
  -d '{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."}'

Response:  
{ "userId": 1, "roles": [ { "name": "Student" } ] }

### PUT /api/go/v1/add-strike-to-user

Adds a strike to a user. If a user reaches 10 strikes, they are banned and a WebSocket ban event is triggered.

curl -X PUT http://localhost:8080/api/go/v1/add-strike-to-user \
  -H "Content-Type: application/json" \
  -d '{"userId": 1}'

Response:  
{ "message": "User's number of strikes successfully implemented" }

### GET /ws

A WebSocket endpoint that broadcasts a real-time event to subscribed clients when a user is banned (i.e., when they reach 10 strikes). To listen:

Use a WebSocket client (e.g., browser or wscat):

wscat -c ws://localhost:8080/ws

When a user reaches 10 strikes, the backend emits:  
{ "userId": 1 }

## 🔐 Working with JWTs locally

1. Sign up with your email and password.  
2. Log in to receive a JWT token.  
3. Use the token for protected endpoints by including it in the Authorization header or request body, depending on implementation.