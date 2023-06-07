package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}

func main() {
	/*
		+	Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis bernilai null / nil
		+	Di Go saat kita membuat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya cth
			- Tipe Data String maka default valuenya adalah string kosong ""
			- Tipe Data Int maka default valuenya adalah 0
			- Tipe Data bool maka default valuenya adalah false
		+ 	Di Go ada data nil, yaitu data kosong
		+ 	Nil sendiri hanya bisa digunakan di beberapa tipe data sepert :
			-	interface
			-	function
			-	map
			-	slice
			-	pointer
			- 	channel
	*/
	helper.PrintBreak("MAP nil")
	var testMap map[string]string = nil
	fmt.Println(testMap)

	helper.PrintBreak("Slice nil")
	var testSlice []string = nil
	fmt.Println(testSlice)

	helper.PrintBreak("Check MAP nil")
	person := NewMap("Rio")
	fmt.Println(person)

	helper.PrintBreak("Check MAP nil2")
	person2 := NewMap("")
	fmt.Println(person2)
	if person2 == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person2)
	}

}
