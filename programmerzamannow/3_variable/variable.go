package main

import "fmt"

func main() {
	var name string

	name = "Rio"
	fmt.Println(name)
	name += "Jozef"
	fmt.Println(name)

	var name2 = "Stefanus"
	fmt.Println(name2)

	name3 := "Antonius"
	fmt.Println(name3)

	var age int
	age = 17
	fmt.Println(age)

	var (
		firstName = "Rio"
		lastName  = "Jozef"
	)

	fmt.Println(firstName, lastName)

	age2, name4 := 18, "Rio Stefanus"
	fmt.Println("Nama = ", name4, ", Umur = ", age2)

}
