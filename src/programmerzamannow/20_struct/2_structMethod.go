package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
	//Married       bool -> jika ini di uncomment cara ke 3 akan error
}

func sayHi(customer Customer, name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (customer Customer) sayHello() {
	fmt.Println("Hello My Name is", customer.Name)
}

func (customer Customer) sayHello2(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	/*
		+ Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
		+ Namun Jika kita ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struct memiliki function
		+ Method adalah function
	*/
	helper.PrintBreak("Struct dengan function biasa")
	var person1 Customer
	person1.Name = "Stefanus"
	sayHi(person1, "Admin")

	helper.PrintBreak("Struct Method")
	person2 := Customer{Name: "Rio"}
	person2.sayHello()

	helper.PrintBreak("Struct Method")
	person3 := Customer{Name: "Jozef"}
	person3.sayHello2("Super Admin")

}
