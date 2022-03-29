package main

import (
	"fmt"
	"strconv"
)

type StringInt int
type StringStr string

type ToString interface {
	ToString() string
}

func (i StringInt) ToString() string {
	return strconv.Itoa(int(i))
}

func (s StringStr) ToString() string {
	return string(s)
}

func stringfy[T ToString](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.ToString())
		switch t := v.(type) {
		default:
			fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
		case StringInt:
			fmt.Printf("StringInt %t\n", t) // t has type bool
		case StringStr:
			fmt.Printf("StringStr %t\n", t) // t has type bool
		}
	}
	return ret
}

func main() {
	fmt.Print(stringfy([]StringInt{1, 2, 3, 4, 5}))
	fmt.Println(stringfy([]StringStr{"1", "2", "3", "4", "5"}))
}
