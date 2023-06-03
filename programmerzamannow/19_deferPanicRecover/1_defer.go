package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func logging() {
	//bukan func logging sesungguhnya cuma perumpamaan
	fmt.Println("Selesai memanggil function")
}

func runApplication(val int) {
	defer logging() //di taruh di awal, error atau tidak tetap akan terpanggil
	fmt.Println("Run Application")
	res := 10 / val
	fmt.Println(res)
}

func main() {
	/*
		+ Defer function adalah function yag bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
		+ Defer function akan selalu dieksekusi walaupun terjadi error di function yang akan dieksekusi
	*/
	helper.PrintBreak("No Error Example")
	runApplication(10) // Tidak Terjadi Error
	helper.PrintBreak("With Error Example")
	runApplication(0) // Terjadi Error

}
