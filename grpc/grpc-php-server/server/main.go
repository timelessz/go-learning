package main

import (
	"godemo/grpc/grpc-php-server/lottery"
	"net"
)

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	grpcServer := grpc.NewServer()
	lottery.RegisterGreeterServer(grpcServer, new(GreeterServerImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

type GreeterServerImpl struct{}

func (p GreeterServerImpl) Lottery(
	ctx context.Context, args *lottery.LotteryReq,
) (*lottery.LotteryRes, error) {
	//reply := &hello.String{Value: "hello:" + args.GetValue()}
	//return reply, nil
	return &lottery.LotteryRes{Data: "hello" + args.GetParam()}, nil
}
