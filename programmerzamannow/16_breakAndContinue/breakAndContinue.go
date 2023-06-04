package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	/*
		Break digunakan untuk menghentikan seluruh perulangan
		Continue digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya
	*/
	helper.PrintBreak("Break Example")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Loops ", i)
	}

	helper.PrintBreak("Continue Example")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
