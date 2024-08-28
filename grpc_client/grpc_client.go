package grpc_client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Yuqizhoujoe/user-service-proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.UserServiceClient
	conn   *grpc.ClientConn
)

func InitGRPCClient() (pb.UserServiceClient, error) {
	if client != nil {
		return client, nil
	}

	// create a context with a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// establish gRPC connection to the UserService using grpc.WithTransportCredentials(insecure.NewCredentials())
	var err error
	conn, err = grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to UserService: %v", err)
	}

	client = pb.NewUserServiceClient(conn)
	return client, nil
}

func CloseGRPCConnection() error {
	if conn != nil {
		return conn.Close()
	}

	return nil
}
