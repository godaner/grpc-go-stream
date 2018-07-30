package student

import (
	"fmt"
	pb "pb"
	"context"
	"io"
)
func GetStudentInfo(client pb.StudentInfoClient, info *pb.StudentInfoRequest) (*pb.StudentInfoResponse) {
	r, err := client.GetStudentInfo(context.Background(),info)
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	return r
}

func UpdateStudentInfo(client pb.StudentInfoClient, notes []*pb.StudentInfoRequest){

	fmt.Println("notes",notes)

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
				fmt.Println("read done ")
				//close(waitReceive)
				waitReceive<-true
				return
			}
			if err != nil {
				fmt.Println("Failed to receive a note : %v", err)
			}
			fmt.Println("Got message %s at point(%d, %d)",in.Id,in.Age,in.Name)
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
