package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"godemo/grpc/grpc-pubsub/pubsubdef"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func (p *PubsubService) Publish(ctx context.Context, arg *pubsubdef.String) (*pubsubdef.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pubsubdef.String{}, nil
}

func (p *PubsubService) SubscribeTopic(arg *pubsubdef.String, stream pubsubdef.PubsubService_SubscribeTopicServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&pubsubdef.String{Value: v.(string)}); nil != err {
			return err
		}
	}
	return nil
}

func (p *PubsubService) Subscribe(arg *pubsubdef.String, stream pubsubdef.PubsubService_SubscribeServer) error {
	ch := p.pub.Subscribe()

	for v := range ch {
		if err := stream.Send(&pubsubdef.String{Value: v.(string)}); nil != err {
			return err
		}
	}
	return nil
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10)}
}

func main() {
	grpcServer := grpc.NewServer()
	pubsubdef.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp4", ":1234")
	if nil != err {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
