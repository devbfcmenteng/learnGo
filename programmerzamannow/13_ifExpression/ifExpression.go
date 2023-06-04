package main

import "fmt"

func main() {
	name := "Rio"

	if name == "Rio" {
		fmt.Println("Welcome Super Admin")
	} else if name == "Admin" {
		fmt.Println("Welcome Home Admin")
	} else {
		fmt.Println("Welcome", name)
	}

	/*
		var length = len(name)
		if length > 5 {
			fmt.Println("Nama Terlalu Panjang")
		} else {
			fmt.Println("Nama Sudah Benar")
		}
		bisa diubah menjadi cth di bawah ini
	*/

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
