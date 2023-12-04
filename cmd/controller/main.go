package main

import (
	"github.com/ragul28/grpc-distributed-worker/internal"
)

func main() {
	internal.GetcontrollerNode().Start()
}
