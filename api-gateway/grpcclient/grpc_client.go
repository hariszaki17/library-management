package grpcclient

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const GrpcTimeout = 5 * time.Second

func NewGrpcConn(host string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithAuthority(host),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: GrpcTimeout,
		}),
	}

	return grpc.NewClient(host, opts...)
}
