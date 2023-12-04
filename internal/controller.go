package internal

import (
	"net"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

// Controller
type ControllerNode struct {
	api     *gin.Engine  // api server
	ln      net.Listener // listener
	svr     *grpc.Server // grpc server
	nodeSvr *Server      // node service
}

func (n *ControllerNode) Init() (err error) {
	// grpc server listener with port as 50051
	n.ln, err = net.Listen("tcp", ":50050")
	if err != nil {
		return err
	}
	return nil
}
