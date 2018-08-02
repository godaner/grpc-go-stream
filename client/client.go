package client

import (
	"log"
	pb "grpc-go-stream/pb"
	"grpc-go-stream/client/student"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func Run() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()



	c := pb.NewStudentInfoClient(conn)


	//GetStudentInfo
	student.GetStudentInfo(c,&pb.StudentInfoRequest{
		Id:1,
		Name:"1231",
		Age:123,
	})


	//UpdateStudentInfo
	student.UpdateStudentInfo(c,[]*pb.StudentInfoRequest{
		{
			Id:1,
			Name:"xiaoyu",
			Age:99,
		},
		{
			Id:2,
			Name:"godaner",
			Age:100,
		},
	})


}