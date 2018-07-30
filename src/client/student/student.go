package student

import (
	"fmt"
	pb "pb"
	"context"
)
func GetStudentInfo(client pb.StudentInfoClient, info *pb.StudentInfoRequest) (*pb.StudentInfoResponse) {
	r, err := client.GetStudentInfo(context.Background(),info)
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	return r
}

func UpdateStudentInfo(client pb.StudentInfoClient, info []*pb.StudentInfoRequest) {

}
