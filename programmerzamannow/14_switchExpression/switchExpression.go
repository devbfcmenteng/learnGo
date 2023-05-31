package main

import "fmt"

func main() {
	username := "Admin"

	switch username {
	case "Admin":
		fmt.Println("Welcome admin")
	case "Superadmin":
		fmt.Println("Welcome admin")
	default:
		if username != "" {
			fmt.Println("Hello", username)
		} else {
			fmt.Println("Hello Guest")
		}

	}

	//Switch dengan Short Statement

	switch length := len(username); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	default:
		fmt.Println("Something missing")
	}

	/*
		Switch without Expression
		Jadi Expressionnya di case nya
	*/

	lengthNew := len(username)
	switch {
	case lengthNew >= 10:
		fmt.Println("Nama Terlalu Panjang")
	case lengthNew >= 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
