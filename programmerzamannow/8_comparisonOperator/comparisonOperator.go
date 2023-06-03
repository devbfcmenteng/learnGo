package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	//Operator Perbandingan
	// >  Lebih Dari
	// <  Kurang Dari
	// >= Lebih Dari Sama Dengan
	// <= Kurang Dari Sama Dengan
	// == Sama Dengan
	// != Tidak Sama Dengan
	// String hanya dapat dibandingkan dengan string, begitu juga Numerik

	helper.PrintBreak("String")
	var name1 = "rio"
	var name2 = "jozef"
	var name3 = "rio"

	var res_name bool = name1 == name2
	var res_name2 bool = name1 == name3

	fmt.Println("name1 = ", name1)
	fmt.Println("name2 = ", name2)
	fmt.Println("name3 = ", name3)
	fmt.Println("name1 == name2 ", res_name)
	fmt.Println("name1 == name3 ", res_name2)

	helper.PrintBreak("Numerik")
	var val1 = 100
	var val2 = 200
	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("val1 > val2 =", val1 > val2)
	fmt.Println("val1 < val2 =", val1 < val2)
	fmt.Println("val1 >= val2 =", val1 >= val2)
	fmt.Println("val1 <= val2 =", val1 <= val2)
	fmt.Println("val1 == val2 =", val1 == val2)
	fmt.Println("val1 != val2 =", val1 != val2)
}
