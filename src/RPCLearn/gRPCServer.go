package RPCLearn

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "grpcTest"
)

func init() {
	fmt.Println("init gRPCServer")
}

const(
	port=":8082"
)

type server struct {}

func (s *server) Add(ctx context.Context,in *pb.AddTarget) (*pb.SumValue, error) {
	return &pb.SumValue{Sum:in.V1+in.V2},nil
}

func (s *server) SayHello(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {
	return &pb.TestReply{Message: "Hello " + in.Name+",this is go grpc"}, nil
}

func TestgRPCServer() {
	fmt.Println("gRPCServer start...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGRPCServer(s,&server{})
	s.Serve(lis)
}
