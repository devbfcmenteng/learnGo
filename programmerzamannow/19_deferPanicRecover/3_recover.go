package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func endApp() {

	message := recover()
	if message != nil {
		fmt.Println("Error Message :", message)
	}
	fmt.Println("Application End")

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error")
	}

	fmt.Println("Application Running")
}

func main() {

	/*
		+ Recover adalah function yang bisa digunakan untuk menangkap data panice seperti catch pada throw exception di bahasa pemrograman lain
		+ Dengan Recover, proses panic akan terhenti, sehingga program akan tetap berjalan
		+ Recover lebih baik di taruh di defer function untuk recover berjalan untuk menankap panic dan program juga tetap berjalan seperti biasa
	*/

	helper.PrintBreak("Example if no error")
	runApp(false)
	fmt.Println("Program tetap berjalan")

	helper.PrintBreak("Example if error")
	runApp(true)
	fmt.Println("Walaupun sudah di panic function Program tetap berjalan karena ada recover, maka message ini akan muncul")

}
