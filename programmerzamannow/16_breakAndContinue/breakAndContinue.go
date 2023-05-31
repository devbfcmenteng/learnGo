package main

import "fmt"

func main() {
	/*
		Break digunakan untuk menghentikan seluruh perulangan
		Continue digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya
	*/

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Loops ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
