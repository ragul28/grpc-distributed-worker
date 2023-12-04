package internal

import (
	proto "github.com/ragul28/grpc-distributed-worker/proto"
)

type Server struct {
	proto.UnimplementedNodeServiceServer

	// channel to receive command
	CmdChannel chan string
}
