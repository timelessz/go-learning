package main

import "fmt"

func main() {
	// 元素类型为 map 的切片: [map, map, map]
	// 先初始化切片, 因为后续需要用索引初始化 map, 所以要给长度
	a := make([]map[string]int, 2)
	// 使用切片的索引对 map 进行初始化
	a[0] = make(map[string]int, 2)
	// 赋值
	a[0]["Tom"] = 100
	a[0]["Tim"] = 99
	// 使用切片的索引对 map 进行初始化
	a[1] = make(map[string]int, 2)
	// 赋值
	a[1]["Tom"] = 100
	a[1]["Tim"] = 99
	fmt.Println(a) // [map[Tim:99 Tom:100] map[]], 这里因为长度为 2, 用零值补齐, 所以会有 map[]

	// 值为切片的 map: map[string][]int, 例: ["语文"][100,99]
	// 先初始化外层的 map
	b := make(map[string][]int, 2)
	// 再初始化内层的切片
	b["语文"] = make([]int, 0, 2)
	// 给切片赋值
	b["语文"] = append(b["语文"], 100, 99)
	b["数学"] = make([]int, 0, 2)
	b["数学"] = append(b["数学"], 97, 93)
	fmt.Println(b)
}
