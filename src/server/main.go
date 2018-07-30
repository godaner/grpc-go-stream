package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "pb"
	"google.golang.org/grpc/reflection"
	"server/student"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterStudentInfoServer(s, &student.StudentInfoServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}