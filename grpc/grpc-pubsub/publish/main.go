package main

import (
	"context"
	"godemo/grpc/grpc-pubsub/pubsubdef"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pubsubdef.NewPubsubServiceClient(conn)
	_, err = client.Publish(
		context.Background(), &pubsubdef.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &pubsubdef.String{Value: "docker: hello Docker"},
	)
	if nil != err {
		log.Fatal(err)
	}
}
