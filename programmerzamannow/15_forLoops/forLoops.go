package main

import "fmt"

func main() {
	counter := 1
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

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Loops for with statement = ", counter2)
	}

	/*
		For bisa digunakan untuk meng
	*/

	fmt.Println("==================Array==================")
	arrayName := [...]string{"Rio", "Stefanus", "Antonius", "Jozef"}
	for i := 0; i < len(arrayName); i++ {
		fmt.Println(arrayName[i])
	}

	fmt.Println("==================Slice==================")
	slice := []string{"Rio", "Stefanus", "Antonius", "Jozef"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("==================Map with for range==================")
	mapPerson := map[string]string{
		"name":    "rio",
		"address": "Tebet",
	}
	fmt.Println(mapPerson)

	/// For Range

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

	for key := range mapPerson { //cara baru
		fmt.Println(mapPerson[key])
	}

	fmt.Println("==================Slice with for range==================")

	for i, val := range slice {
		fmt.Println("Index ", i, "=", val)
	}

	for _, val := range slice {
		fmt.Println(val)
	}

	// underscore diatas digunakan untuk menunjukan bahwa variable tersebut tidak dibutuhkan

}
