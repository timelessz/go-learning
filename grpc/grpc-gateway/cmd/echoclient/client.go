package main

import (
	"context"
	pb "godemo/grpc/grpc-gateway/echopb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	addr = "localhost:8088"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	log.Printf("echo request: wang\n")

	r, err := c.Echo(ctx, &pb.StringMessage{Value: "wang"})
	if err != nil {
		log.Fatalf("could not echo: %v\n", err)
	}

	log.Printf("Echo reply: %s\n", r.GetValue())
}
