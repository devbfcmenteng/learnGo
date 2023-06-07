package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {

	alamat1 := new(Address) // new(type) -> membuat pointer kosong

	alamat1.Country = "Indonesia"
	fmt.Println(alamat1)
}
