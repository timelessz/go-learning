package main

import "fmt"

type Userdemo struct {
	Name string
	Age  byte
}

func main() {
	d := [...]Userdemo{
		{"TigerwolfC", 20}, // 可省略元素类型。
		{"chen_peggy", 18}, // 别忘了最后一行的逗号。
	}
	fmt.Println(d)
}
