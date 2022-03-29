package main

import (
	"fmt"
	"godemo/gokit/v1/calculate"
	"net/http"
)

func main() {
	server := calculate.NewService()
	endpoints := calculate.NewEndPointServer(server)
	httpHandler := calculate.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
