package student

import (
	pb "pb"
	"context"
	"io"
	"fmt"
)
type StudentInfoServerImpl struct{
	StudentInfoResponse	[]*pb.StudentInfoResponse
}

func (server *StudentInfoServerImpl) GetStudentInfo(ctx context.Context, in *pb.StudentInfoRequest) (*pb.StudentInfoResponse, error){
	return &pb.StudentInfoResponse{
		Id:987,
		Name:"张大可",
		Age:9998,
	}, nil
}

func (server *StudentInfoServerImpl) UpdateStudentInfo(stream pb.StudentInfo_UpdateStudentInfoServer)(error){
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			return nil
		}
		if err != nil {
			fmt.Println("ERR",err)
			return err
		}
		fmt.Println("userinfo ",in)
		for _, r := range server.StudentInfoResponse{
			if err := stream.Send(r); err != nil {
				return err
			}
		}
	}
}