package main

import (
	"context"
	"log"
	"net"

	pb "../../proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement user.UserService
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUsers implements user.UserService
func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	// Static user data
	users := []*pb.User{
		{
			Id:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Location: "New York",
		},
		{
			Id:       2,
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Location: "San Francisco",
		},
		{
			Id:       3,
			Name:     "Bob Johnson",
			Email:    "bob@example.com",
			Location: "Chicago",
		},
	}

	return &pb.GetUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
