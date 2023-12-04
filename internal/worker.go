package internal

import (
	proto "github.com/ragul28/grpc-distributed-worker/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerNode struct {
	conn *grpc.ClientConn        // grpc client connection
	c    proto.NodeServiceClient // grpc client
}

func (n *WorkerNode) Init() (err error) {
	// connect to controller node
	n.conn, err = grpc.Dial(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// grpc client
	n.c = proto.NewNodeServiceClient(n.conn)

	return nil
}
