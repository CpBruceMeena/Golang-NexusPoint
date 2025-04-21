# Golang NexusPoint

A microservices-based application demonstrating gRPC and HTTP communication between services.

## Project Structure

```
.
├── golang-central/     # Central Service (gRPC + HTTP Server)
│   ├── main.go        # Main server implementation
│   ├── product_service.go  # Product service implementation
│   └── go.mod         # Go module definition
│
├── golang-grpc-app/   # Client Service (gRPC Client + HTTP Server)
│   ├── main.go        # Client implementation
│   └── go.mod         # Go module definition
│
├── golang-json-app/   # Simple HTTP Service
│   ├── main.go        # HTTP server implementation
│   └── go.mod         # Go module definition
│
└── proto/             # Centralized Protocol Buffers
    ├── user/          # User service definitions
    │   └── v1/
    │       └── user.proto
    ├── product/       # Product service definitions
    │   └── v1/
    │       └── product.proto
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

# Terminal 3 - Start JSON Service
cd golang-json-app
go run main.go
```

## Services

### Central Service (gRPC + HTTP Server)
- Runs on:
  - HTTP port 8080
  - gRPC port 50051
- Implements:
  - HTTP endpoints:
    - GET /get-users: Returns list of users
    - GET /get-profile?user_id=<id>: Returns user profile
  - gRPC services:
    - User Service:
      - GetUsers: Returns list of users
      - GetProfile: Returns user profile by ID
    - Product Service:
      - GetProducts: Returns list of products
      - GetProduct: Returns product by ID

### Client Service (HTTP Server + gRPC Client)
- Runs on port 8082
- Exposes HTTP endpoints:
  - GET /users: Returns list of users
  - GET /profile?user_id=<id>: Returns user profile
  - GET /products: Returns list of products
  - GET /product?id=<id>: Returns product by ID

### JSON Service (Simple HTTP Server)
- Runs on port 8081
- Demonstrates a simple HTTP-only microservice
- Exposes HTTP endpoints:
  - GET /users: Returns list of users in JSON format
  - GET /profile?user_id=<id>: Returns user profile in JSON format
- Features:
  - Simple HTTP server implementation
  - Direct JSON response handling
  - No gRPC dependencies
  - Easy to understand and modify
  - Good starting point for new developers

## API Endpoints

### Central Service HTTP Endpoints
```bash
# Get all users
curl http://localhost:8080/get-users

# Get user profile
curl http://localhost:8080/get-profile?user_id=1
```

### Client Service HTTP Endpoints
```bash
# Get all users
curl http://localhost:8082/users

# Get user profile
curl http://localhost:8082/profile?user_id=1

# Get all products
curl http://localhost:8082/products

# Get product by ID
curl http://localhost:8082/product?id=1
```

### JSON Service HTTP Endpoints
```bash
# Get all users
curl http://localhost:8081/users

# Get user profile
curl http://localhost:8081/profile?user_id=1
```

## Adding a New Microservice

To add a new microservice (e.g., Order Service):

1. Create new proto definitions:
```bash
mkdir -p proto/order/v1
touch proto/order/v1/order.proto
```

2. Define the service in the proto file:
```protobuf
syntax = "proto3";

package order.v1;

option go_package = "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/order/v1";

// Define your messages and service
message Order {
  int32 id = 1;
  int32 user_id = 2;
  repeated int32 product_ids = 3;
  float total = 4;
}

service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse) {}
  rpc GetOrder(GetOrderRequest) returns (GetOrderResponse) {}
}
```

3. Generate proto files:
```bash
cd proto
./generate.sh
```

4. Implement the service in golang-central:
```bash
touch golang-central/order_service.go
```

5. Add service implementation:
```go
package main

import (
    "context"
    orderv1 "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/order/v1"
)

type orderServer struct {
    orderv1.UnimplementedOrderServiceServer
}

// Implement service methods
```

6. Register the service in main.go:
```go
func main() {
    // ...
    orderv1.RegisterOrderServiceServer(s, &orderServer{})
    // ...
}
```

7. Add HTTP endpoints in golang-central:
```go
// Add new HTTP handlers in golang-central/main.go
http.HandleFunc("/get-orders", getOrdersHandler)
http.HandleFunc("/get-order", getOrderHandler)
```

8. Update the client service to expose HTTP endpoints:
```go
// Add new HTTP handlers in golang-grpc-app/main.go
```

## Development Guidelines

1. **Proto Definitions**:
   - Keep all proto files in the centralized `proto` directory
   - Use versioning (v1, v2, etc.) for API changes
   - Follow proto3 syntax and best practices

2. **Service Implementation**:
   - Each service should have its own implementation file
   - Keep business logic separate from transport layer
   - Use proper error handling and status codes
   - Implement both gRPC and HTTP endpoints in golang-central
   - For simple services, consider using golang-json-app as a template

3. **Code Generation**:
   - Always run `./generate.sh` after modifying proto files
   - Generated code is in `proto/gen/go/`
   - Never modify generated files directly

4. **Testing**:
   - Add unit tests for service implementations
   - Test both gRPC and HTTP endpoints
   - Use mock services for integration tests

5. **Version Control**:
   - Commit proto changes and generated code together
   - Use descriptive commit messages
   - Follow semantic versioning for API changes

## Best Practices

1. **API Design**:
   - Use meaningful message and field names
   - Document all RPC methods and messages
   - Consider backward compatibility
   - Keep HTTP and gRPC APIs consistent
   - For simple services, prefer HTTP for easier debugging

2. **Error Handling**:
   - Use appropriate gRPC status codes
   - Provide meaningful error messages
   - Handle edge cases gracefully
   - Return consistent error formats for both HTTP and gRPC

3. **Performance**:
   - Use connection pooling for gRPC clients
   - Implement proper timeouts
   - Consider streaming for large datasets
   - Cache frequently accessed data
   - For simple services, HTTP might be sufficient

4. **Security**:
   - Validate all input data
   - Implement proper authentication
   - Use TLS for production deployments
   - Apply rate limiting