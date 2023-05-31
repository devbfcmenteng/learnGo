package main

import "fmt"

func getHello(firstName string, lastName string, age string) string {
	return "Hello " + firstName + " " + lastName + ", age = " + age
}

func getHello2(name string) string {
	if name != "" {
		return "Hello " + name
	} else {
		return "Hello Guest"
	}

}

func main() {
	res1 := getHello("Rio", "Jozef", "22")
	fmt.Println(res1)
	firstname := "Stefanus"
	lastname := "Antonius"
	age := "17"
	fmt.Println(getHello(firstname, lastname, age))

	fmt.Println("=================================")
	res2 := getHello2("Rio")
	fmt.Println(res2)
	fmt.Println(getHello2("Admin"))
	fmt.Println(getHello2(""))
}
