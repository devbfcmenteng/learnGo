package main

import "fmt"

func convertString(num byte) string {
	str := string(num)
	return str
}

func main() {
	fmt.Println("Rio")
	fmt.Println("Rio Jozef")
	//	function untuk String
	//	len("string") => Menghitung jumlah karakter string
	//		cth : len("Rio") => 3
	fmt.Println(len("Rio"))
	// "String"[number] => Mengambil karakter pada posisi yang ditentukan
	//		cth : "Rio"[0] => 83 -> 83 disini adalah byte dari String 'R'	fmt.Println("Rio"[0])
	fmt.Println("Rio"[0])
	fmt.Println(string("Rio"[0]))
	fmt.Println(convertString("Rio"[0]))
}
