package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	/*
		+ Function adalah first class citizen
		+ Fcuntion juga merupakan tipe data, dan bisa disimpan di dalam variable, bisa di buat secara indipendent tanpa harus membuat group2 function / class
	*/
	goodbye := getGoodBye    //variable goodbye yang memuat function
	result := goodbye("Rio") //memanggil func dari variabel goodbye
	fmt.Println(result)

}
