package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	productv1 "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/product/v1"
	user_pb "github.com/CpBruceMeena/golang-nexuspoint/proto/gen/go/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	user_pb.UnimplementedUserServiceServer
}

// GetUsers implements user.UserService
func (s *server) GetUsers(ctx context.Context, req *user_pb.GetUsersRequest) (*user_pb.GetUsersResponse, error) {
	users := getStaticUsers()

	// Convert to gRPC format
	grpcUsers := make([]*user_pb.User, len(users))
	for i, u := range users {
		grpcUsers[i] = &user_pb.User{
			Id:       int32(u.ID),
			Name:     u.Name,
			Email:    u.Email,
			Location: u.Location,
		}
	}

	return &user_pb.GetUsersResponse{Users: grpcUsers}, nil
}

// GetProfile implements user.UserService
func (s *server) GetProfile(ctx context.Context, req *user_pb.GetProfileRequest) (*user_pb.GetProfileResponse, error) {
	// Static profile data
	profiles := map[int32]*user_pb.Profile{
		1: {
			Id:      1,
			Bio:     "Software Engineer with 5 years of experience",
			Website: "https://johndoe.com",
			Company: "Tech Corp",
			Role:    "Senior Developer",
		},
		2: {
			Id:      2,
			Bio:     "Product Manager passionate about user experience",
			Website: "https://janesmith.com",
			Company: "Innovate Inc",
			Role:    "Product Lead",
		},
		3: {
			Id:      3,
			Bio:     "DevOps Engineer specializing in cloud infrastructure",
			Website: "https://bobjohnson.com",
			Company: "Cloud Solutions",
			Role:    "DevOps Lead",
		},
	}

	profile, exists := profiles[req.UserId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "profile not found for user %d", req.UserId)
	}

	return &user_pb.GetProfileResponse{Profile: profile}, nil
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
		user_pb.RegisterUserServiceServer(s, &server{})
		productv1.RegisterProductServiceServer(s, &productServer{})
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
