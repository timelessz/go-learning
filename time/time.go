package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间
	now := time.Now()
	fmt.Println(now)
	//时间戳，单位秒
	seconds_stamp := now.Unix()
	fmt.Println(seconds_stamp)
	//时间戳，单位纳秒 1秒=1e3毫秒=1e9纳秒
	fmt.Println(now.UnixNano())
	//时间戳，单位毫秒
	fmt.Println(now.UnixNano() / 1e6)
	//只输出年，其他字段类似
	fmt.Println(now.Year())
	//只输出月
	fmt.Println(now.Month())
	//只输出天
	fmt.Println(now.Day())
	//只输小时
	fmt.Println(now.Hour())
	//一年的第几天
	fmt.Println(now.YearDay())
	//年月日
	fmt.Println(now.Date())
	//时分秒
	fmt.Println(now.Clock())
}
