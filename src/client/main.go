package main

import (
	"log"
	"google.golang.org/grpc"
	pb "pb"
	"client/student"
	"fmt"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()



	c := pb.NewStudentInfoClient(conn)


	//GetStudentInfo
	s := student.GetStudentInfo(c,&pb.StudentInfoRequest{
		Id:1,
		Name:"1231",
		Age:123,
	})
	fmt.Println("userinfo is : ",s.GetAge()," , ",s.GetAge()," , ",s.GetName())

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