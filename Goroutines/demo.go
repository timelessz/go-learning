package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5) //通道，返回结果数据
	done := make(chan bool)   //通道
	go func() {
		for {
			j, ok := <-jobs
			//fmt.Printf("ok的类型为%T\n",ok) //ok的类型为bool 当通道关闭ok的类型为false
			if ok {
				fmt.Println("收到工作", j)
			} else {
				fmt.Println("收到全部工作结果")
				done <- true //其实这里放true和false都无所谓
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		time.Sleep(time.Second)
		jobs <- j
		fmt.Println("sent  job", j)
	}
	close(jobs)
	fmt.Println("发送完毕")
	//等待工作
	<-done
}
