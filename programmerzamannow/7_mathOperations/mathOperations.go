package main

import "fmt"

func main() {
	// + -> Tambah
	// - -> Kurang
	// * -> Perkalian
	// / -> Pembagian
	// % -> Sisa Pembagian

	var a = 10
	var b = 10
	var plus = a + b
	var mul = a * b
	fmt.Println(plus)
	fmt.Println(mul)

	fmt.Println("==============")
	fmt.Println("Augmented Assignment")

	//Augmented Assignments
	// a = a + 10 -> a += 10
	// a = a - 10 -> a -= 10
	// a = a * 10 -> a *= 10
	// a = a / 10 -> a /= 10
	// a = a % 10 -> a %= 10

	var a_aug = 10

	fmt.Println("Val for Augmented = ", a_aug)

	a_aug += 10
	fmt.Println("Plus 10 Augmented = ", a_aug)

	a_aug -= 5
	fmt.Println("Minus 5 Augmented = ", a_aug)

	a_aug *= 2
	fmt.Println("Mulltiply 2 Augmented = ", a_aug)

	a_aug /= 3
	fmt.Println("Division 3 Augmented = ", a_aug)

	a_aug %= 2
	fmt.Println("Mod 2 Augmented = ", a_aug)

	fmt.Println("==============")
	fmt.Println("Unary Operator")

	//Unary Operator
	// i++ -> i = i + 1
	// i-- -> i = i - 1
	// -i -> tetap sama bilangan negatif i
	// +i -> tetap sama bilangan positif i
	// !i -> Boolean kebalikan misal i = true -> !i nilai maka menjadi false

	var i = 1
	fmt.Println("i =", i)

	i++
	fmt.Println("i++ =", i)

	i--
	fmt.Println("i-- =", i)

	var bol bool = true
	fmt.Println("var bol =", bol)
	fmt.Println("var !bol =", !bol)
}
