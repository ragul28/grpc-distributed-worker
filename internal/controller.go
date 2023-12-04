package internal

import (
	"net"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

// Controller
type ControllerNode struct {
	api     *gin.Engine
	ln      net.Listener
	svr     *grpc.Server
	nodeSvr *Server
}

func (n *ControllerNode) Init() (err error) {
	n.ln, err = net.Listen("tcp", ":50050")
	if err != nil {
		return err
	}
	return nil
}

func (n *ControllerNode) Start() {
	// start grpc server
	go n.svr.Serve(n.ln)

	// start api server
	_ = n.api.Run(":9092")

	// wait for exit
	n.svr.Stop()
}

var controllerNode *ControllerNode

// GetcontrollerNode returns the node instance
func GetcontrollerNode() *ControllerNode {
	if controllerNode == nil {
		// node
		controllerNode = &ControllerNode{}

		// initialize node
		if err := controllerNode.Init(); err != nil {
			panic(err)
		}
	}

	return controllerNode
}
