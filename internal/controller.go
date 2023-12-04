package internal

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	proto "github.com/ragul28/grpc-distributed-worker/proto"

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

	// grpc server
	n.svr = grpc.NewServer()

	// node service
	n.nodeSvr = GetGrpcServer()

	// register node service to grpc server
	proto.RegisterNodeServiceServer(n.svr, n.nodeSvr)

	// api
	n.api = gin.Default()
	n.api.POST("/tasks", func(c *gin.Context) {
		// parse payload
		var payload struct {
			Cmd string `json:"cmd"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// send command to node service
		n.nodeSvr.CmdChannel <- payload.Cmd

		c.AbortWithStatus(http.StatusOK)
	})

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
