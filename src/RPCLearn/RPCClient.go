package RPCLearn

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func TestRPCClient() {

	//向1号服务请求
	go startRpcClient2()
	//创建服务
	go startRpcServer2()

	strInput := ""
	fmt.Scanln(&strInput)
}

func startRpcServer2() {
	//start a server
	fmt.Println("Start RPC Service 2....")
	rpc.RegisterName("HelloClient",new(HelloClient))
	listener,err:=net.Listen("tcp",":8091")
	if err!=nil{
		log.Fatal("LisenTCP error:",err)
	}

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Fatal("Accept error:",err)
		}
		go rpc.ServeConn(conn)
	}
}

type HelloClient struct {}

func (p *HelloClient)Hello(request string,reply *string) error {
	*reply="from 2: "+request
	return nil
}

//RPC Client
func startRpcClient2() {
	client,err:=rpc.Dial("tcp","localhost:8090")
	if err!=nil{
		log.Fatal("dialing err:",err)
	}
	var reply string
	err=client.Call("HelloService.Hello","this is 2",&reply)
	if err!=nil{
		log.Fatal("Call err:",err)
	}
	fmt.Println(reply)
}