package main

import (
	"context"
	"fmt"
	"godemo/grpc/grpc-hello/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.String{Value: "baidu.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
