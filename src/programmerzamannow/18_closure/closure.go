package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	/*
		+ Closure adalah kemampuann sebuah function untuk berinteraksi dengan data-data disekitarnya dalam scope yang sama data scope diatasnya
		+ Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi, karena kita dapat secara tidak sengaja mengubah data yang tidak kita inginkan
	*/
	name := "Rio"
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		name = "Jozef"
		counter++ //karena di anonymous func ini tidak di declarasikan variable counter, maka variable counter scope yang diatasnya yang akan dipakai dan berubah, maka harus berhati2 dalam melakukan code seperti ini, begitu juga dengan variable name
		fmt.Println("name var in func = ", name)
		fmt.Println("counter var in func =", counter)

		// jika seperti dibawah ini name dan counter di deklarasikan, maka name dan counter variabel diatas (di scope sebelum function ini) tidak akan berubah nilainya
		/*
			name := "Jozef"
			counter := 0
			counter++
			fmt.Println("name var in func = ", name)
			fmt.Println("counter var in func =", counter)
		*/

	}

	increment()
	increment()

	helper.PrintBreak("")
	fmt.Println("counter first var =", counter)
	fmt.Println("name first var =", name)

}
