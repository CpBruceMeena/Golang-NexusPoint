# Golang Central Service

A simple Go HTTP service that exposes a `/get-users` endpoint returning static user data.

## Requirements

- Go 1.21 or higher

## Running the Service

1. Navigate to the project directory:
   ```bash
   cd golang-central
   ```

2. Run the service:
   ```bash
   go run main.go
   ```

The service will start on port 8080.

## Available Endpoints

### GET /get-users

Returns a static list of users in JSON format.

Example request:
```bash
curl http://localhost:8080/get-users
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