package main

import (
	"encoding/json"
	"os"
)

type Address struct {
	City string
}

type VCard struct {
	Name     string
	Addresss []Address
}

func main() {
	pa := Address{"jinan"}
	wa := Address{"beijing"}
	vc := VCard{"ZX", []Address{pa, wa}}
	file, _ := os.OpenFile("zx.json", os.O_CREATE|os.O_RDWR, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.Encode(vc)
}
