package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Rio"
	names[1] = "Stefanus"
	names[2] = "Antonius"
	names[3] = "Jozef"

	//bisa juga di tulis seperti ini
	/*
		var names = [4]string{
			"Rio",
			"Stefanus",
			"Antonius",
			"Jozef",
		}
	*/

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	fmt.Println("================")

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("================")

	var values = [3]int{
		100,
		90,
		80,
	}
	fmt.Println(values)

	fmt.Println("================")

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}
	values[2] = 85
	fmt.Println(values)

	//Keterangan
	/*
		len(array) 				-> Untuk mendapatkan panjang Array
		array[index]			-> Mendapatkan data di posisi index
		array[index] = value	-> Merubah data di posisi index
	*/

}
