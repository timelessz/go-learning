package main

import (
	"fmt"
	"time"
)

func main() {
	//1、时间戳转时间
	nowUnix := time.Now().Unix() //获取当前时间戳
	nowStr := unixToStr(nowUnix, "2006-01-02 15:04:05")
	fmt.Printf("1、时间戳转时间：%d => %s \n", nowUnix, nowStr)
	//2、时间转时间戳
	nowStr = time.Now().Format("2006/01/02 15:04:05") //根据指定的模板[ 2006/01/02 15:04:05 ]，返回时间。
	nowUnix, err := strToUnix(nowStr, "2006/01/02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("2、时间转时间戳：%s => %d", nowStr, nowUnix)
}

//时间戳转时间
func unixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

//时间转时间戳
func strToUnix(timeStr, layout string) (int64, error) {
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
