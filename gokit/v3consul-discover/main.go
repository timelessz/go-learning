package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"godemo/gokit/v2-discover/discover"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建环境变量
	//var (
	//	consulHost = flag.String("consul.host", "", "consul server ip address")
	//	consulPort = flag.String("consul.port", "", "consul server port")
	//)
	//flag.Parse()

	consulHost := localIP()
	consulPort := "8500"

	//创建日志组件
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//创建consul客户端对象
	var client consul.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = "http://" + consulHost + ":" + consulPort
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}

	ctx := context.Background()
	//创建Endpoint
	discoverEndpoint := discover.MakeDiscoverEndpoint(ctx, client, logger)
	//创建传输层
	r := discover.MakeHttpHandler(discoverEndpoint)
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		logger.Log("transport", "HTTP", "addr", "9001")
		errc <- http.ListenAndServe(":9001", r)
	}()

	// 开始运行，等待结束
	logger.Log("exit", <-errc)
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
