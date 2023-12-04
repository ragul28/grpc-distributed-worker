package internal

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

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

func (n *WorkerNode) Start() {
	// log
	fmt.Println("worker node started")

	// report status
	_, _ = n.c.ReportStatus(context.Background(), &proto.Request{})

	// assign task
	stream, _ := n.c.AssignTask(context.Background(), &proto.Request{})
	for {
		// receive command from controller node
		res, err := stream.Recv()
		if err != nil {
			return
		}

		// log command
		fmt.Println("received command: ", res.Data)

		// execute command
		parts := strings.Split(res.Data, " ")
		if err := exec.Command(parts[0], parts[1:]...).Run(); err != nil {
			fmt.Println(err)
		}
	}
}

var workerNode *WorkerNode

func GetWorkerNode() *WorkerNode {
	if workerNode == nil {
		// node
		workerNode = &WorkerNode{}

		// initialize node
		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
