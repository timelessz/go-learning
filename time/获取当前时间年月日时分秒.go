package main

import (
	"fmt"
	"time"
)

func main() {
	//golang的time.Format方法一定要用2006-01-02 15:04:05作参数？
	// 1、年月日
	year := time.Now().Year()
	month := time.Now().Month()
	//或者
	//month := time.Now().Month().String()
	day := time.Now().Day()

	//2、年月日
	year1 := time.Now().Year()
	month1 := time.Now().Format("01")
	day1 := time.Now().Day()

	//3、年月日，时分秒
	year2 := time.Now().Format("2006")
	month2 := time.Now().Format("01")
	day2 := time.Now().Format("02")
	hour := time.Now().Format("15")
	min := time.Now().Format("04")
	second := time.Now().Format("05")

	fmt.Println("1、年月日---------------------")
	fmt.Printf("%d, %s, %d \n", year, month, day)
	fmt.Println("2、年月日---------------------")
	fmt.Printf("%d, %s, %d \n", year1, month1, day1)
	fmt.Println("3、年月日，时分秒---------------------")
	fmt.Printf("%s, %s, %s, %s, %s, %s", year2, month2, day2, hour, min, second)
	fmt.Println("")
}
