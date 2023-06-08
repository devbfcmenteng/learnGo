package main

import (
	"bfcTest/packages/helper" //mengimport package
	"fmt"
)

func main() {
	/*
		+ Di bahasa pemrograman lain, biasanya ada kata kunci yag bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
		+ Di Go-Lang untuk menentukan access modifier cukup dengan nama function atau variable
		+ Jika nama func atau variable diawali dengan huruf besar, maka bisa diakses dari package lain, jika dimulai dengan huruf kecil berarti tidak bisa diakses package lain
	*/
	helper.PrintBreak("")           //cth function yang bisa diakses dari package helper karena berawalkan hurut besar
	fmt.Println(helper.Application) //cth variable yang bisa diakses dari package helper karena berawalkan hurut besar

}
