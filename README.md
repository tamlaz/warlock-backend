# Warlock Backend

A backend service written in Go for managing the server-side tasks for Warlock AI.  
My main purpose with this project is to get familiar with Go and its ecosystem.

## 📦 Dependencies

See `go.mod` for the full list. Notable ones:  
`github.com/gin-gonic/gin`, `gorm.io/gorm`, `gorm.io/driver/postgres`, `github.com/gorilla/websocket`

## 🛠️ Installation

```bash
git clone https://github.com/tamlaz/warlock-backend.git  
cd warlock-backend  
go mod download  
go run main.go
```

## Running the app as a Docker container
```bash
cd warlock-backend
docker compose -f docker-compose-go.yml up -d
```

## 📡 API Endpoints

All routes assume the server is running locally on http://localhost:8080.

| Method | Path                                  | Purpose                               | Auth required |
|--------|---------------------------------------|---------------------------------------|---------------|
| GET    | /health                               | Health-check (“Is backend alive?”)    | ❌            |
| POST   | /api/go/v1/signup                     | Register a new user                   | ❌            |
| POST   | /api/go/v1/login                      | Authenticate user & receive JWT       | ❌            |
| POST   | /api/go/v1/validate-user              | Validate a JWT and fetch user info    | ❌            |
| PUT    | /api/go/v1/add-strike-to-user         | Increment a user's strike count       | ❌            |
| POST   | /api/go/v1/save-qa                    | Save a new Q&A message pair           | ❌            |
| GET    | /api/go/v1/get-conversation-history   | Retrieve AI message history for user  | ❌            |
| GET    | /ws                                   | Subscribe to real-time ban events     | ❌            |

### GET /health

Performs a simple health check.

```bash
curl http://localhost:8080/health
```

Response:

```json
{"Warlock backend is up and running"}
```

### POST /api/go/v1/signup

Registers a new user.

```bash
curl -X POST http://localhost:8080/api/go/v1/signup \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "Str0ngP@ssw0rd!", "firstName": "Alice", "lastName": "Wonder"}'
```

Response:

```json
{ "message": "User created successfully" }
```

### POST /api/go/v1/login

Authenticates a user and returns a signed JWT.

```bash
curl -X POST http://localhost:8080/api/go/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "Str0ngP@ssw0rd!"}'
```

Response:

```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..." }
```

### POST /api/go/v1/validate-user

Validates a JWT and fetches user information.

```bash
curl -X POST http://localhost:8080/api/go/v1/validate-user \
  -H "Content-Type: application/json" \
  -d '{"warlock_api_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."}'
```

Response:

```json
{ "user_id": 1, "roles": ["Student"] }
```

### PUT /api/go/v1/add-strike-to-user

Adds a strike to a user. If a user reaches 10 strikes, they are banned and a WebSocket ban event is triggered.

```bash
curl -X PUT http://localhost:8080/api/go/v1/add-strike-to-user \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1}'
```

Response:

```json
{"User's number of strikes successfully implemented" }
```

### POST /api/go/v1/save-qa

Saves a Q&A message pair associated with a user.

```bash
curl -X POST http://localhost:8080/api/go/v1/save-qa \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "topic_id": 5, "subject_id": 3, "human_message_content": "What is Go?", "ai_message_content": "Go is a statically typed language."}'
```

Response:

```json
{ "message": "New QA record saved successfully" }
```

### GET /api/go/v1/get-conversation-history

Retrieves AI message history for a specific user.

```bash
curl "http://localhost:8080/api/go/v1/get-conversation-history?user_id=1"
```

Response:

```json
[
  { "message_content": "Go is a statically typed language.", "message_type": "AI" },
  { "message_content": "Go has excellent concurrency support.", "message_type": "AI" }
]
```

### GET /ws

A WebSocket endpoint that broadcasts a real-time event when a user is banned (upon reaching 10 strikes).

Use a WebSocket client to connect:

```bash
wscat -c ws://localhost:8080/ws
```

When a user reaches 10 strikes, the backend emits:

```json
{ "userId": 1 }
```

## 🔐 Working with JWTs locally

1. Sign up with your email and password.  
2. Log in to receive a JWT token.  
3. Use the token for protected endpoints by including it in the Authorization header or request body, depending on implementation.
