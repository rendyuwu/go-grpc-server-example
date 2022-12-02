package grpc

import (
	"context"
	pb "github.com/rendyuwu/go-grpc-server-example/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServer
	Users []*pb.UserResponse
}

func NewUserService() *UserService {
	var users []*pb.UserResponse

	// dummy data for example
	users = append(users, &pb.UserResponse{
		Name:     "Rendy Wijaya",
		Username: "rendyuwu",
		Age:      22,
	})
	users = append(users, &pb.UserResponse{
		Name:     "Budi Setiawan",
		Username: "budi",
		Age:      25,
	})

	return &UserService{Users: users}
}

func (u *UserService) FindUserByEmail(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	for _, user := range u.Users {
		if user.Username == in.Username {
			return user, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User %s not found", in.Username)
}
