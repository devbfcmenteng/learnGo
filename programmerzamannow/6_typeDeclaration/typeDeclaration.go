package main

import "fmt"

func main() {
	type NoKTP string
	//membuat tipe baru 'NoKTP' dengan tipe string

	var ktpRio NoKTP = "12345678"
	//jadi var ktpRio NoKTP = "12345678" itu sama degan
	// var ktpRio string ="12345678"

	type Married bool
	var marriedStatus Married = false

	fmt.Println(ktpRio)
	fmt.Println(marriedStatus)

}
