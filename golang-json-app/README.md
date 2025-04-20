# Golang JSON App

A Go service that fetches and serves user data from the golang-central service.

## Requirements

- Go 1.21 or higher
- The golang-central service must be running on port 8080

## Running the Service

1. First, ensure the golang-central service is running:
   ```bash
   cd ../golang-central
   go run main.go
   ```

2. In a new terminal, navigate to this directory:
   ```bash
   cd golang-json-app
   ```

3. Run the service:
   ```bash
   go run main.go
   ```

The service will start on port 8081.

## Available Endpoints

### GET /users

Fetches user data from the golang-central service and returns it in JSON format.

Example request:
```bash
curl http://localhost:8081/users
```

Example response:
```json
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "location": "New York"
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@example.com",
        "location": "San Francisco"
    },
    {
        "id": 3,
        "name": "Bob Johnson",
        "email": "bob@example.com",
        "location": "Chicago"
    }
]
``` 