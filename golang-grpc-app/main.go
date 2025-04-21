package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	user_pb "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = ":8082"
)

var (
	grpcClient user_pb.UserServiceClient
)

// User represents a user in JSON format
type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

// Profile represents a profile in JSON format
type Profile struct {
	ID      int32  `json:"id"`
	Bio     string `json:"bio"`
	Website string `json:"website"`
	Company string `json:"company"`
	Role    string `json:"role"`
}

// initGrpcClient initializes the gRPC client
func initGrpcClient() error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	grpcClient = user_pb.NewUserServiceClient(conn)
	return nil
}

// getUsersHandler handles HTTP requests and makes gRPC call to central service
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Make gRPC call
	resp, err := grpcClient.GetUsers(context.Background(), &user_pb.GetUsersRequest{})
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

// getProfileHandler handles HTTP requests for profile and makes gRPC call to central service
func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user_id from query parameter
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	// Make gRPC call
	resp, err := grpcClient.GetProfile(context.Background(), &user_pb.GetProfileRequest{UserId: int32(userID)})
	if err != nil {
		log.Printf("Failed to get profile: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Convert gRPC response to JSON format
	profile := Profile{
		ID:      resp.Profile.Id,
		Bio:     resp.Profile.Bio,
		Website: resp.Profile.Website,
		Company: resp.Profile.Company,
		Role:    resp.Profile.Role,
	}

	// Set response headers and encode JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func main() {
	// Initialize gRPC client
	if err := initGrpcClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Start HTTP server
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/profile", getProfileHandler)
	log.Printf("HTTP server starting on port %s...", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}
