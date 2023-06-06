package main

import (
	"bfcTest/packages/helper"
	"fmt"
	"reflect"
)

func Ups(i int) interface{} { //interface{} -> untuk memungkinkan return parameter tipe datanya bebas
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func getDetailData(data interface{}) (string, interface{}) { // checkTypeVale(val interface{}) -> untuk memungkinkan parameter yang masuk tipe datanya bebas
	typeVal := reflect.TypeOf(data).Kind().String()
	return typeVal, data

}

func main() {
	/*
		+	Go-Lang Bukan bahasa pemrograman berorientasi objek
		+ 	Biasaya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap
			sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
			Cth di Java ada java.lang.Object
		+	Untuk mengangan kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
		+	Interface kosong adalah interface yg tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

		Ada banyak contoh penggunaan interface kosong di Go :
		+	fmt.Println(a ...interface[])
		+	panic(v interface)
		+ 	recover() interface
	*/

	helper.PrintBreak("Example 1")
	// var data1 int = Ups(1) -> ini akan error sekalipun func Ups ada yang mereturn int
	var data1 interface{} = Ups(1)
	fmt.Println(data1)

	helper.PrintBreak("Example 2")
	var data2 interface{} = Ups(2)
	fmt.Println(data2)

	helper.PrintBreak("Example 3")
	var data3 interface{} = Ups(3)
	fmt.Println(data3)

	helper.PrintBreak("Example 4")
	data4 := Ups(4)
	fmt.Println(data4)

	helper.PrintBreak("Example 5")
	dataVal1 := "Rio"
	type1, val1 := getDetailData(dataVal1)
	fmt.Printf("Data = %v \nType Data = %v \nData Value = %v \n", dataVal1, type1, val1)

	helper.PrintBreak("Example 6")
	dataVal2 := 20
	type2, val2 := getDetailData(dataVal2)
	fmt.Printf("Data = %v \nType Data = %v \nData Value = %v \n", dataVal2, type2, val2)

	helper.PrintBreak("Example 7")
	dataVal3 := true
	type3, val3 := getDetailData(dataVal3)
	fmt.Printf("Data = %v \nType Data = %v \nData Value = %v \n", dataVal3, type3, val3)
}
