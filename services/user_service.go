package services

import (
	"context"
	"log"
	"travelgo/grpc_client"
	"travelgo/models"

	pb "github.com/Yuqizhoujoe/user-service-proto/proto"
)

type UserService struct {
	grpcClient pb.UserServiceClient
}

func NewUserService() (*UserService, error) {
	client, err := grpc_client.InitGRPCClient()
	if err != nil {
		return nil, err
	}

	return &UserService{
		grpcClient: client,
	}, nil
}

func (us *UserService) AddUser(data models.AddUser) (models.AddUserResponse, error) {
	ctx := context.Background()

	// trigger gRPC AddUser
	addUserReq := pb.AddUserRequest{
		Email: data.Email,
	}
	_, err := us.grpcClient.AddUser(ctx, &addUserReq)
	if err != nil {
		log.Panicf("Failed to call AddUser: %v", err)
		return models.AddUserResponse{
			Success: false,
		}, err
	}

	return models.AddUserResponse{
		Success: true,
	}, nil
}
