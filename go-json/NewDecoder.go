package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City string
	}
	type VCard struct {
		Name     string
		Addresss []Address
	}
	file, _ := os.OpenFile("zx.json", os.O_CREATE|os.O_RDWR, 0)
	defer file.Close()
	dec := json.NewDecoder(file)
	vc := &VCard{}
	dec.Decode(vc)
	fmt.Printf("%v", vc)
}
