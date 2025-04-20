package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	pb "github.com/CpBruceMeena/golang-nexuspoint/golang-central/proto"

	"google.golang.org/grpc"
)

const (
	httpPort = ":8080"
	grpcPort = ":50051"
)

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

// server is used to implement user.UserService
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUsers implements user.UserService
func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users := getStaticUsers()

	// Convert to gRPC format
	grpcUsers := make([]*pb.User, len(users))
	for i, u := range users {
		grpcUsers[i] = &pb.User{
			Id:       int32(u.ID),
			Name:     u.Name,
			Email:    u.Email,
			Location: u.Location,
		}
	}

	return &pb.GetUsersResponse{Users: grpcUsers}, nil
}

// getStaticUsers returns the static list of users
func getStaticUsers() []User {
	return []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Location: "New York"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Location: "San Francisco"},
		{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", Location: "Chicago"},
	}
}

// getUsersHandler handles HTTP requests
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := getStaticUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, &server{})
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server
	http.HandleFunc("/get-users", getUsersHandler)
	log.Printf("HTTP server starting on port %s...", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}
