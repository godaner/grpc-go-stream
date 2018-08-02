package student

import (
	"fmt"
	pb "grpc-go-stream/pb"
	"context"
	"io"
)
func GetStudentInfo(client pb.StudentInfoClient, info *pb.StudentInfoRequest) {
	r, err := client.GetStudentInfo(context.Background(),info)
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	fmt.Println("from server , GetStudentInfo userinfo is : ",r.GetAge()," , ",r.GetAge()," , ",r.GetName())
}

func UpdateStudentInfo(client pb.StudentInfoClient, notes []*pb.StudentInfoRequest){

	fmt.Println("to server , UpdateStudentInfo notes is : ",notes)

	stream, err := client.UpdateStudentInfo(context.Background())
	if err != nil {
		fmt.Println("%v.RouteChat(_) = _, %v", client, err)
	}
	waitReceive := make(chan bool)

	// read from server
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				//read done.
				//fmt.Println("read done ")
				//close(waitReceive)
				waitReceive<-true
				return
			}
			if err != nil {
				fmt.Println("Failed to receive a note : %v", err)
			}

			fmt.Println("from server , UpdateStudentInfo info is : ",in.Id,in.Age,in.Name)
		}
	}()
	// send to server
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			fmt.Println("Failed to send a note: %v", err)
		}
	}

	stream.CloseSend()
	<-waitReceive
}
