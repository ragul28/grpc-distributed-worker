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
