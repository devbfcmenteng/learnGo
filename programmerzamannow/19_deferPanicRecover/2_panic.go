package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func endApp() {
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
		+ Panic function adalah function yag bisa kita gunakan untuk menghentikan program seperti throw exception di bahasa pemrograman yang lain
		+ Panic biasanya dipanggil ketika terjadi error pada saat program kita berjalan
		+ Saat Panic function dipanggil, program akan terhenti, namun Defer function tetap akan dieksekusi
	*/

	helper.PrintBreak("Example if no error")
	runApp(false)
	fmt.Println("Program tetap berjalan")

	helper.PrintBreak("Example if error")
	runApp(true)
	fmt.Println("Program berhenti di 'panic' ini tidak akan muncul")

}
