package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	var num32 int32 = 100000
	var num64 int64 = int64(num32)
	var num8 int8 = int8(num32)

	fmt.Println("num32 = ", num32)
	fmt.Println("num64 = ", num64)
	fmt.Println("num8 = ", num8)

	helper.PrintBreak("")

	var num32_2 int32 = 128
	var num64_2 int64 = int64(num32_2)
	var num8_2 int8 = int8(num32_2)

	fmt.Println("num32_2 = ", num32_2)
	fmt.Println("num64_2 = ", num64_2)
	fmt.Println("num8_2 = ", num8_2)

	helper.PrintBreak("")

	var num32_3 int32 = 129
	var num64_3 int64 = int64(num32_3)
	var num8_3 int8 = int8(num32_3)

	fmt.Println("num32_3 = ", num32_3)
	fmt.Println("num64_3 = ", num64_3)
	fmt.Println("num8_3 = ", num8_3)

	helper.PrintBreak("")

	var name string = "Rio"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println("Name = ", name)
	fmt.Println("Get first char from", name, " = ", e)
	fmt.Println("Convert byte", e, "to String = ", eString)

}
