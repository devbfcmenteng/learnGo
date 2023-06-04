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

func main() {
	/*
		+	Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya didalam satu kesatua
		+	Struct biasanya representasi data dalam program aplikasi yang kita buat
		+	Data di Struct disimpan didalam field
		+	Struct adalah kumpulan dari field
		+ 	Struct adalah template data atau prototype data
		+ 	Struct tidak bisa langsung digunakan
		+	Namun kita bisa membuat data/object dari strcut yang telah kita buat
	*/
	helper.PrintBreak("Struct Cara1")
	var person1 Customer
	person1.Name = "Rio"
	person1.Address = "Tebet"
	person1.Age = 34

	fmt.Println(person1)

	fmt.Println(person1.Name)
	fmt.Println(person1.Address)
	fmt.Println(person1.Age)

	/*
		Struct Literals
		- Langkah diatas adalah salah satu cara membuat struct, namun ada banyak cara yg bisa kita gunakan untuk membuat data dr struct
	*/

	helper.PrintBreak("Struct Cara2")
	person2 := Customer{
		Name:    "Stefanus",
		Address: "Kelingkit",
		Age:     30,
	}
	fmt.Println(person2)

	fmt.Println(person2.Name)
	fmt.Println(person2.Address)
	fmt.Println(person2.Age)

	helper.PrintBreak("Struct Cara3")
	person3 := Customer{"Antonius", "sapta", 25} //Urutan sesuai dengan struct yang telah dibuat diatas, ini cara yang paling tidak baik,karena jika struct ada penambahan field maka akan error, cara 1 dan 2 tidak akan error dianjurkan untuk tidak digunakan
	fmt.Println(person3)

	fmt.Println(person3.Name)
	fmt.Println(person3.Address)
	fmt.Println(person3.Age)

}
