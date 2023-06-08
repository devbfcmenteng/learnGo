package helper

import "fmt"

func PrintBreak(title string) {
	switch title != "" {
	case true:
		fmt.Println("===============", title, "===============")
	default:
		fmt.Println("=========================================")
	}
}

var verion string = "1.0.0" //diawali huruf kecil tidak dapat diakses dari luar
var Application string = "Aplikasi belajar golang basic"

func helloWorld() { //diawali huruf kecil tidak dapat diakses dari luar
	fmt.Println("Hello")
}
