package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func getFullName() (string, string, string) { //Video 23
	return "Rio", "Stefanus", "Jozef"
}

func getFullName2() (firstName, middleName, lastName string) { //Named Return Values Video 24
	//	func getFullName2() (firstName string, middleName string, lastName string) { -> bisa seperti ini, biasanya seperti ini kalau returnnya valuenya berbeda2 ada int string, dll
	firstName = "Monkey"
	middleName = "D"
	lastName = "Luffy"
	/*
		untuk return bisa 2 cara
		- return -> hanya return saja karena variable sudah di declare di func
		- return firstName, middleName, lastName -> tidak efisien jika return variable banyak
	*/
	return

}

func main() {

	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
	firstName2, _, lastName2 := getFullName() //Ignore return second val
	fmt.Println(firstName2, lastName2)
	firstName3, middleName3, _ := getFullName() //Ignore return third val
	fmt.Println(firstName3, middleName3)
	firstName4, _, _ := getFullName() //Ignore return second and third val
	fmt.Println(firstName4)
	_, _, lastName5 := getFullName() //Ignore return first and second val
	fmt.Println(lastName5)
	/*
		can't ignore all return value, it will error, example :
		_, _, _ := getFullName()

	*/
	helper.PrintBreak("Named Return Value")
	firstName_1, middleName_1, lastName_1 := getFullName2()
	fmt.Println(firstName_1, middleName_1, lastName_1)
}
