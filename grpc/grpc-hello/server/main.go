package main

import (
	"context"
	"godemo/grpc/grpc-hello/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(
		grpcServer,
		new(HelloServiceImpl),
	)
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *hello.String,
) (*hello.String, error) {
	reply := &hello.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
