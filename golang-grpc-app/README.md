# Golang gRPC App

A gRPC service that serves user data.

## Requirements

- Go 1.21 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Generate gRPC code:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/user.proto
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on port 50051.

## Testing the Service

You can test the service using `grpcurl`:

```bash
# Install grpcurl if you haven't already
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List available services
grpcurl -plaintext localhost:50051 list

# Call the GetUsers method
grpcurl -plaintext -d '{}' localhost:50051 user.UserService/GetUsers
```

## Available Endpoints

### GetUsers

Returns a list of users in gRPC format.

Example response:
```json
{
  "users": [
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
}
``` 