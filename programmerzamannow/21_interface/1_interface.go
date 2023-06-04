package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	/*
		+ interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
		+ Sebuah interface berisikan definisi-definisi method
		+ Biasanya interface digunakan sebagai kontrak

		+ Setiap tipe data yang sesuai dengan kontark interface, secara otomatis dianggap sebagai interface tersebut, sehinggak kita tidak perlu mengimplementasi interface secara manual
		+ Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface yang mana

	*/
	helper.PrintBreak("Example 1")
	var person1 Person
	person1.Name = "Rio"

	SayHello(person1)

	helper.PrintBreak("Example 2")
	animal := Animal{Name: "Kucing"}
	SayHello(animal)

}
