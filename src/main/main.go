package main

import (
	"RPCLearn"
	_ "RPCLearn"
	"fmt"
)

func init() {
	fmt.Println("init main")
}

func main() {
	//RPC Server
	RPCLearn.TestRPCServer()
	//RPC Client
	RPCLearn.TestRPCClient()

	//gRPC Server
	RPCLearn.TestgRPCServer()
}
