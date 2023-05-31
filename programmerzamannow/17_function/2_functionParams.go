package main

import "fmt"

func hello(firstName string, lastName string, age int) {
	fmt.Println("Hello", firstName, lastName, ", age = ", age)
}

func main() {
	hello("Rio", "Jozef", 12)
	firstname := "Stefanus"
	lastname := "Antonius"
	age := 17
	hello(firstname, lastname, age)
}
