package handlers

import (
	"context"

	pb "godemo/gokit/All"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.EchoServer {
	return echoService{}
}

type echoService struct{}

func (s echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp.Out = "dsadas"
	return &resp, nil
}

func (s echoService) Louder(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	return &resp, nil
}
