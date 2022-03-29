package main

import (
	"context"
	"fmt"
	hellostream "godemo/grpc/grpc-stream/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hellostream.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hellostream.String{Value: "baidu.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&hellostream.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
