# Golang NexusPoint

A microservices-based application demonstrating gRPC communication between services.

## Project Structure

```
.
├── golang-central/     # Central Service (gRPC Server)
├── golang-grpc-app/    # Client Service (gRPC Client + HTTP Server)
└── proto/             # Centralized Protocol Buffers
    ├── user/
    │   └── v1/
    │       └── user.proto
    ├── gen/           # Generated code
    │   └── go/
    ├── generate.sh    # Script to generate proto files
    ├── buf.gen.yaml   # Buf generation configuration
    └── buf.yaml       # Buf workspace configuration
```

## Prerequisites

- Go 1.22.0 or higher
- Buf CLI (for protocol buffer generation)

## Setup

1. Install Buf CLI:
```bash
go install github.com/bufbuild/buf/cmd/buf@latest
```

2. Generate Protocol Buffer files:
```bash
cd proto
chmod +x generate.sh
./generate.sh
```

3. Start the services:
```bash
# Terminal 1 - Start Central Service
cd golang-central
go run main.go

# Terminal 2 - Start Client Service
cd golang-grpc-app
go run main.go
```

## Services

### Central Service (gRPC Server)
- Runs on port 50051
- Implements user service with:
  - GetUsers: Returns list of users
  - GetProfile: Returns user profile by ID

### Client Service (HTTP Server + gRPC Client)
- Runs on port 8082
- Exposes HTTP endpoints:
  - GET /users: Returns list of users
  - GET /profile?user_id=<id>: Returns user profile

## API Endpoints

### Get Users
```bash
curl http://localhost:8082/users
```

### Get Profile
```bash
curl http://localhost:8082/profile?user_id=1
```

## Development

- Protocol Buffer definitions are centralized in the `proto` directory
- Use `./generate.sh` in the proto directory to regenerate files after changes
- Services import proto definitions from the centralized location