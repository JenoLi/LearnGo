package RPCLearn

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func init() {
	fmt.Println("init RPCServer")
}

func TestRPCServer() {
	//创建1号服务
	go startRpcServer1()

	strinput:=""
	fmt.Scanln(&strinput)
	if(strinput=="a"){
		//向2号服务请求
		startRpcClient1()
	}
}

func startRpcServer1() {
	fmt.Println("Start RPC Service 1....")

	rpc.RegisterName("HelloService",new(HelloService))
	listener,err:=net.Listen("tcp",":8090")
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
type HelloService struct {}

func (p *HelloService)Hello(request string,reply *string) error {
	*reply="from 1:"+request
	return nil
}


func startRpcClient1() {
	client,err:=rpc.Dial("tcp","localhost:8091")
	if err!=nil{
		log.Fatal("dialing err:",err)
	}
	var reply string
	err=client.Call("HelloClient.Hello","this is 1",&reply)
	if err!=nil{
		log.Fatal("Call err:",err)
	}
	fmt.Println(reply)
}