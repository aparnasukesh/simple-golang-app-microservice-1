// package grpcclient

// import (
// 	"microservice-one/grpcproto"

// 	"google.golang.org/grpc"
// )

// func NewMicroServiceTwoServiceClient(port string) (grpcproto.MicroServiceTwoServiceClient, error) {
// 	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return grpcproto.NewMicroServiceTwoServiceClient(conn), nil
// }

package grpcclient

import (
	"fmt"
	"microservice-one/grpcproto"
	"os"

	"google.golang.org/grpc"
)

func NewMicroServiceTwoServiceClient() (grpcproto.MicroServiceTwoServiceClient, error) {
	host := os.Getenv("MICROSERVICE_TWO_HOST")
	port := os.Getenv("MICROSERVICE_TWO_PORT")

	if host == "" || port == "" {
		return nil, fmt.Errorf("MICROSERVICE_TWO_HOST or MICROSERVICE_TWO_PORT not set")
	}

	// Use Kubernetes service DNS
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return grpcproto.NewMicroServiceTwoServiceClient(conn), nil
}
