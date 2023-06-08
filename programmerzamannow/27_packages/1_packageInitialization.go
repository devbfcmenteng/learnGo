package main

import (
	"bfcTest/packages/database" //akan langsung mengeksekui function init()
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	/*
		+ 	Saat kita membuat package, kita bisa membuat sebuah function yang akan dieksekusi ketika package kita diakases
		+ 	Ini sangan cocok ketika contohnya, jika package kita berisi function2 untuk berkomunikasi dengan databse,
			kita bisa membuat function inisialisasi untuk membuat koneksi ke database
		+	Untuk membuat function yang diakses secara otomatsi ketika package diakses, kita cukup membuat function dengan nama init
	*/
	helper.PrintBreak("") //cth function yang bisa diakses dari package helper karena berawalkan hurut besar
	database := database.GetDatabase()
	fmt.Println("Database", database)
}
