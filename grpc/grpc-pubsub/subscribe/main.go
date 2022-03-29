package main

import (
	"context"
	"fmt"
	"godemo/grpc/grpc-pubsub/pubsubdef"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if nil != err {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pubsubdef.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &pubsubdef.String{Value: "golang:"},
	)
	if nil != err {
		log.Fatal(err)
	}
	go func() {
		for {
			reply, err := stream.Recv()
			if nil != err {
				if io.EOF == err {
					break
				}
				log.Fatal(err)
			}
			fmt.Println("sub1: ", reply.GetValue())
		}
	}()
	streamTopic, err := client.SubscribeTopic(
		context.Background(), &pubsubdef.String{Value: "golang:"},
	)
	if nil != err {
		log.Fatal(err)
	}
	go func() {
		for {
			reply, err := streamTopic.Recv()
			if nil != err {
				if io.EOF == err {
					break
				}
				log.Fatal(err)
			}
			fmt.Println("subTopic: ", reply.GetValue())
		}
	}()
	<-make(chan bool)
}
