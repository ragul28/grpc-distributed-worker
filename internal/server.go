package internal

import (
	"context"

	proto "github.com/ragul28/grpc-distributed-worker/proto"
)

type Server struct {
	proto.UnimplementedNodeServiceServer

	// channel to receive command
	CmdChannel chan string
}

func (n Server) ReportStatus(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	return &proto.Response{Data: "ok"}, nil
}

func (n Server) AssignTask(request *proto.Request, server proto.NodeService_AssignTaskServer) error {
	for {
		select {
		case cmd := <-n.CmdChannel:
			// receive command and send to worker node (client)
			if err := server.Send(&proto.Response{Data: cmd}); err != nil {
				return err
			}
		}
	}
}

var server *Server

// GetNodeServiceGrpcServer singleton service
func GetGrpcServer() *Server {
	if server == nil {
		server = &Server{
			CmdChannel: make(chan string),
		}
	}
	return server
}
