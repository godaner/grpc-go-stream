package student

import (
	pb "pb"
	"context"
	"io"
	"fmt"
)
type StudentInfoServerImpl struct{
}

func (server *StudentInfoServerImpl) GetStudentInfo(ctx context.Context, in *pb.StudentInfoRequest) (*pb.StudentInfoResponse, error){

	fmt.Println("from client , GetStudentInfo stream is : ",in)

	return &pb.StudentInfoResponse{
		Id:in.Id+1,
		Name:"server back : "+in.Name,
		Age:in.Age+1,
	}, nil
}

func (server *StudentInfoServerImpl) UpdateStudentInfo(stream pb.StudentInfo_UpdateStudentInfoServer)(error){
	for {
		// read info from client
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			break
		}
		if err != nil {
			fmt.Println("ERR",err)
			return err
		}
		fmt.Println("from client , UpdateStudentInfo stream is : ",in)

	}
	//send info to client

	StudentInfoResponse:=[]*pb.StudentInfoResponse{
		{
			Id:1,
			Name:"zk",
			Age:222,
		},
	}
	for _, r := range StudentInfoResponse{
		if err := stream.Send(r); err != nil {
			return err
		}
	}
	return nil
}