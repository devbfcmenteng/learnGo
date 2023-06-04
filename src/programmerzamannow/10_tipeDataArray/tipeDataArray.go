package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	var names [4]string
	names[0] = "Rio"
	names[1] = "Stefanus"
	names[2] = "Antonius"
	names[3] = "Jozef"

	//bisa juga di tulis seperti ini
	/*
		var names = [4]string{
			"Rio",
			"Stefanus",
			"Antonius",
			"Jozef",
		}
	*/
	helper.PrintBreak("Print with key index")
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	helper.PrintBreak("Print with for Loop")

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	var values = [3]int{
		100,
		90,
		80,
	}
	helper.PrintBreak("New Array for test")
	fmt.Println(values)

	helper.PrintBreak("Print with for Loop")

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	helper.PrintBreak("Change Last Array val to 85")
	values[2] = 85
	fmt.Println(values)

	//Keterangan
	/*
		len(array) 				-> Untuk mendapatkan panjang Array
		array[index]			-> Mendapatkan data di posisi index
		array[index] = value	-> Merubah data di posisi index
	*/

}
