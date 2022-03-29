package main

import (
	"context"
	"fmt"
	"godemo/grpc/grpc-php-server/lottery"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := lottery.NewGreeterClient(conn)
	reply, err := client.Lottery(context.Background(), &lottery.LotteryReq{
		Param: "dsadsa",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetData())
}
