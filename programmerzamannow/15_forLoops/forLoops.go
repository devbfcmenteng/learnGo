package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	counter := 1
	helper.PrintBreak("For Loops")
	for counter <= 10 {
		fmt.Println("Loops = ", counter)
		counter++
	}

	/*
		For dengan Statement

		- Dalam for kita bisa menambahkan statement, dimana terdaoat 2 statement yang bisa ditambahkan di for
		- Init Statement, yaitu statement sebelum for di eksekusi
		- Post Statement, yaitu statement yang akan selalu diekseskusi diakhir tiap perulangan
	*/

	helper.PrintBreak("For Loops with statement")
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Loops for with statement = ", counter2)
	}

	/*
		For bisa digunakan untuk meng
	*/

	helper.PrintBreak("Array")
	arrayName := [...]string{"Rio", "Stefanus", "Antonius", "Jozef"}

	helper.PrintBreak("Array for loop")
	for i := 0; i < len(arrayName); i++ {
		fmt.Println(arrayName[i])
	}

	helper.PrintBreak("Slice")
	slice := []string{"Rio", "Stefanus", "Antonius", "Jozef"}

	helper.PrintBreak("Slice for loops")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	helper.PrintBreak("Map")
	mapPerson := map[string]string{
		"name":    "rio",
		"address": "Tebet",
	}
	fmt.Println(mapPerson)

	/// For Range
	helper.PrintBreak("Map with for range")
	for key, value := range mapPerson {
		fmt.Println("key =", key, ",value = ", value)
	}

	/*
		Ignore variable yang hanya berlaku di for tetapi tidak berlaku di function :

		Jika ingin meng-ignore value pada mapPerson maka dapat dilakukan
		for key, _ := range mapPerson {
		jika ingin disingkat maka dapat menjadi
		for key := range mapPerson {

		tetapi jika ingin meng-ignore key seperti dibawah ini
		for _, value := range mapPerson {
		maka tidak dapat disingkat seperti key diatas
		for value := range mapPerson { -> disini tetap akan me-return key

	*/
	helper.PrintBreak("Map with for range ignore value")
	for key := range mapPerson { //cara baru
		fmt.Println(mapPerson[key])
	}

	helper.PrintBreak("Slice with for range")
	for i, val := range slice {
		fmt.Println("Index ", i, "=", val)
	}

	helper.PrintBreak("Slice with for range ignore index / key")
	for _, val := range slice {
		fmt.Println(val)
	}

	// underscore diatas digunakan untuk menunjukan bahwa variable tersebut tidak dibutuhkan

}
