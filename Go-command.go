package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var domains = []string{
		"baidu.com",
		"qiangbi.net",
		"google.com",
		"alibaba.com",
		"studygolang.com",
	}
	for _, domain := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			cmd := exec.Command("dig", "-t", "mx", "+short", domain)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("获取异常！")
			}
			fmt.Printf("\n%s\n", string(out))
		}(domain)
	}
	wg.Wait()
}
