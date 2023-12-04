package main

import (
	"os"

	"github.com/ragul28/grpc-distributed-worker/internal"
)

func main() {
	nodeType := os.Args[1]
	switch nodeType {
	case "controller":
		internal.GetcontrollerNode().Start()
	case "worker":
		internal.GetWorkerNode().Start()
	default:
		panic("invalid node type")
	}
}
