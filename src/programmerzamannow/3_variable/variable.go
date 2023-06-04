package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	var name string

	name = "Rio"
	fmt.Println(name)
	name += "Jozef"
	fmt.Println(name)
	helper.PrintBreak("")

	var name2 = "Stefanus"
	fmt.Println(name2)

	helper.PrintBreak("")
	name3 := "Antonius"
	fmt.Println(name3)

	helper.PrintBreak("")
	var age int
	age = 17
	fmt.Println(age)

	helper.PrintBreak("")
	var (
		firstName = "Rio"
		lastName  = "Jozef"
	)

	fmt.Println(firstName, lastName)

	helper.PrintBreak("")
	age2, name4 := 18, "Rio Stefanus"
	fmt.Println("Nama = ", name4, ", Umur = ", age2)

}
