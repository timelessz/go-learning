package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type UserInfo struct {
	UserId   int
	UserName string
}

func main() {
	traceId := uuid.New()
	userId := 1111
	ctx := context.WithValue(context.Background(), "trace_id", traceId.String())
	ctx = context.WithValue(ctx, "user_id", userId)
	go getUserInfo(ctx)
	time.Sleep(time.Second * 2)
	fmt.Printf("\n[main] (trace_id: %v, user_id: %v)",
		ctx.Value("trace_id").(string), ctx.Value("user_id").(int))
}

func getUserInfo(ctx context.Context) UserInfo {
	fmt.Printf("[getUserInfo] (trace_id: %v, user_id: %v) get user info",
		ctx.Value("trace_id").(string), ctx.Value("user_id").(int))
	return UserInfo{UserId: ctx.Value("user_id").(int), UserName: "Tome"}
}
