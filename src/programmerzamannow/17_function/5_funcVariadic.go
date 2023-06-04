package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, val := range numbers {
		total += val
	}

	return total
}

func testIseng(names ...string) (fullname string) {
	length := len(names)
	for i, name := range names {
		if length-1 == i {
			fullname += name
		} else {
			fullname += name + " "
		}

	}

	return
}

func main() {
	/*
		+ Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
		+ Varargs artinya datanya bisa menerima lebih dari satu input atau anggap saja semacam Array.
		+ Apa bedanya dengan paratameter biasa dengen tipe data Array?
			-	Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
			- 	Jika  menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma
	*/

	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	slice := []int{
		1, 2, 3, 4, 5,
	}

	totalSlice := sumAll(slice...)
	fmt.Println(totalSlice)

	names := testIseng("Rio", "Stefanus", "Antonius", "Jozef")
	fmt.Println(names)

}
