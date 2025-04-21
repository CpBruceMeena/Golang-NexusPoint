package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/CpBruceMeena/golang-nexuspoint/golang-grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = ":8082"
)

// User represents a user in JSON format
type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

// getUsersHandler handles HTTP requests and makes gRPC call to central service
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create gRPC connection to central service
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Create gRPC client
	client := pb.NewUserServiceClient(conn)

	// Make gRPC call
	resp, err := client.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Convert gRPC response to JSON format
	users := make([]User, len(resp.Users))
	for i, u := range resp.Users {
		users[i] = User{
			ID:       u.Id,
			Name:     u.Name,
			Email:    u.Email,
			Location: u.Location,
		}
	}

	// Set response headers and encode JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Start HTTP server
	http.HandleFunc("/users", getUsersHandler)
	log.Printf("HTTP server starting on port %s...", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}
