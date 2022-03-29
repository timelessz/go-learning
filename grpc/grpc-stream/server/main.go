package main

import (
	"context"
	hellostream "godemo/grpc/grpc-stream/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	hellostream.RegisterHelloServiceServer(
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
	ctx context.Context, args *hellostream.String,
) (*hellostream.String, error) {
	reply := &hellostream.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream hellostream.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &hellostream.String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
