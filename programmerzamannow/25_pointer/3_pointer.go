package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Makasar"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	helper.PrintBreak("")
	*address2 = Address{"Malang", "Jawa Timur", "indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
