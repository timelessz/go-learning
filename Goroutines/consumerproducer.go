package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com",
		"http://www.github.com/",
		"https://golang.google.cn/",
	}
	for _, url := range urls {
		// 增加WaitGroup的数量
		wg.Add(1)
		// 分发一个goroutine请求对应的url
		go func(url string) {
			// 当这个goroutine完成的时候，减少WaitGroup的数量
			defer wg.Done()
			// Fetch the URL.
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res.StatusCode)
		}(url)
	}
	// 等待至所有的请求都完成
	wg.Wait()
}
