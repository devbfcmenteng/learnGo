package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {
	/*
		+ Secara default di go lang semua variable di passing by value, bukan referance
		+ artinya jika kita mengirim sebuah variable ke dalam function, method, atau variable lain sebenarnya yang dikirim adalah duplikasi valuenya
	*/
	var str1 string = "Rio"
	var str2 *string = &str1

	helper.PrintBreak("str1")
	fmt.Printf("str1 (Value str1) = %v, \n&str1 (memory address str1)= %v, \n*str1 (tidak ada karena bukan pointer) \n", str1, &str1)

	helper.PrintBreak("str2")
	fmt.Printf("var str2 *string = &str1 \nstr2 (memory address dr str1) = %v, \n&str2 (memory address str2)= %v, \n*str2 (Value str2) = %v \n", str2, &str2, *str2)

	helper.PrintBreak("Contoh dengan stuct")
	//var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//var address2 *Address = &address1
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
}
