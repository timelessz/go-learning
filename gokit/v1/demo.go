package main

import "fmt"

type TestStruct struct {
	id   int
	name string
}

type AnotherTestStruct struct {
	id   int
	name string
}

type ITest interface {
	SayHello()
}

func main() {
	var itest ITest
	test := TestStruct{id: 1, name: "test1"}
	test2 := AnotherTestStruct{id: 2, name: "anotherTest"}
	itest = test     // OK
	itest.SayHello() // test1

	itest = &test    // OK
	itest.SayHello() // test1

	//itest = test2    // cannot use test2 (type AnotherTestStruct) as type ITest in assignment:
	//// AnotherTestStruct does not implement ITest (SayHello method has pointer receiver)
	//itest.SayHello()

	itest = &test2   // OK
	itest.SayHello() // anotherTest
}

func (test TestStruct) SayHello() {
	fmt.Println(test.name)
}

func (test *AnotherTestStruct) SayHello() {
	fmt.Println("dsdsadsa" + test.name)
}
