package main

import (
	"bfcTest/packages/helper"
	"fmt"
	"reflect"
)

func random() interface{} {
	return "OK"
}

func random2() interface{} {
	return 100
}

func random3(val interface{}) interface{} {
	return val
}

func printRes(val interface{}) {
	switch value := val.(type) {
	case string:
		fmt.Println("Value =", value, "is string")
	case int:
		fmt.Println("Value =", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}
}

func main() {
	/*
		Type Asserstions
		+ Type Assertion merupakan kemampuan merubah tipe data yang diinginkan
		+ Sering digunakan ketika kita bertemu dengan data interface kosong
	*/
	helper.PrintBreak("Example 1 Data")
	res := random()
	fmt.Printf("Res = %v, Data Type = %v \n", res, reflect.TypeOf(res).Kind().String())

	helper.PrintBreak("Example 1 Data after Type Assertion")
	resultString := res.(string)
	fmt.Printf("Res = %v, Data Type = %v \n", resultString, reflect.TypeOf(resultString).Kind().String())

	/*
		Cth dibawah ini akan terjadi error, karena function interface diatas return string dan kita merubah dengan type assertion ke int

			resultInt := res.(int)
			fmt.Println(resultInt)


		Saat salah menggunakan type assertion maka bisa berakibat terjadi panic di aplikasi
		Jika Panic tidak ter cover maka otomatis program akan mati
		agar lebih aman sebaiknya kita menggunaka swtich expression untuk melakukan type assertion seperti contoh dibawah ini
	*/
	helper.PrintBreak("Example 2 With Switch Case")

	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value =", value, "is string")
	case int:
		fmt.Println("Value =", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

	helper.PrintBreak("Example 2 With Switch Case")

	var result2 interface{} = random2()

	switch value := result2.(type) {
	case string:
		fmt.Println("Value =", value, "is string")
	case int:
		fmt.Println("Value =", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

	helper.PrintBreak("Example 3 With Selfmade function")
	printRes(random3("test"))
	helper.PrintBreak("Example 3 With Selfmade function")
	printRes(random3(12))

}
