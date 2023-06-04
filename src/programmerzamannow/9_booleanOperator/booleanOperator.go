package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	//Boolean Operator Perbandingan
	// && -> Dan
	//		Untuk Operator &&
	//		val1 = true,	val2 = true,	Result = true
	//		val1 = true, 	val2 = false,	Result = false
	//		val1 = false, 	val2 = true,	Result = false
	//		val1 = false, 	val2 = false,	Result = false
	// ||  -> Atau
	//		Untuk Operator ||
	//		val1 = true,	val2 = true,	Result = true
	//		val1 = true, 	val2 = false,	Result = true
	//		val1 = false, 	val2 = true,	Result = true
	//		val1 = false, 	val2 = false,	Result = false
	// ! Kebalikan
	//		Untuk Operator !
	//		val = true,		Result = false
	//		val = false,	Result = true

	var val1 bool = true
	var val2 bool = true

	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("val1 && val2", val1 && val2)
	fmt.Println("val1 || val2", val1 || val2)

	helper.PrintBreak("")

	val2 = false
	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("val1 && val2", val1 && val2)
	fmt.Println("val1 || val2", val1 || val2)

	helper.PrintBreak("")

	val1 = false
	val2 = true
	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("val1 && val2", val1 && val2)
	fmt.Println("val1 || val2", val1 || val2)

	helper.PrintBreak("")

	val2 = false
	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("val1 && val2", val1 && val2)
	fmt.Println("val1 || val2", val1 || val2)

	helper.PrintBreak("")

	val1 = true
	fmt.Println("val1 =", val1)
	fmt.Println("val2 =", val2)
	fmt.Println("!val1 = ", !val1)
	fmt.Println("!val2 =", !val2)

	//Cth lagi
	var nilaiAkhir = 90
	var absensi = 80
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("lulus? = ", lulus)
}
