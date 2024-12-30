package grpcclient

import (
	"microservice-one/grpcproto"

	"google.golang.org/grpc"
)

func NewMicroServiceTwoServiceClient(port string) (grpcproto.MicroServiceTwoServiceClient, error) {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return grpcproto.NewMicroServiceTwoServiceClient(conn), nil
}
