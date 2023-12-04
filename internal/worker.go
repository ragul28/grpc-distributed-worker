package internal

import (
	proto "github.com/ragul28/grpc-distributed-worker/proto"

	"google.golang.org/grpc"
)

type WorkerNode struct {
	conn *grpc.ClientConn        // grpc client connection
	c    proto.NodeServiceClient // grpc client
}
