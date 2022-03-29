package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"godemo/grpc/protoc/pb/proto_demo"
)

func main() {
	test := &proto_demo.Student{
		Name:    "James",
		Male:    true,
		Scores:  []int32{98, 85, 88},
		Subject: map[string]int32{"age": 18, "level": 1},
	}
	// 序列化
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("proto encode error: ", err)
		return
	}
	// 反序列化
	newTest := &proto_demo.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("proto decode error: ", err)
	}
	if test.GetScores()[1] != newTest.GetScores()[1] {
		fmt.Printf("data mismatch score %d != %d", test.GetScores()[1], newTest.GetScores()[1])
		return
	}
	if test.GetName() != newTest.GetName() {
		fmt.Printf("data mismatch name %s != %s", test.GetName(), newTest.GetName())
		return
	}
	fmt.Println("data match!")
}
