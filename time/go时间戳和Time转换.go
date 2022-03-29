package main

import (
	"fmt"
	"time"
)

func main() {
	//Time类型.Unix  是将Time类型转为时间戳
	timestamp := time.Now().Unix() //time.Now()是当前时间（Time类型）
	fmt.Println("now", timestamp)
	//time.Unix  是time包里的函数，将时间戳转为Time类型
	fmt.Println(time.Unix(timestamp, 0))
	// 时间戳 转 Time
	nowStr := "2013/7/11"
	nowUnix, _ := strToUnixD(nowStr, "2006/1/02")
	print(nowUnix)
	fmt.Println(time.Unix(nowUnix, 0))
}

//时间转时间戳
func strToUnixD(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}
