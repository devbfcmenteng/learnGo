package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

type Man struct {
	Name    string
	Address string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	helper.PrintBreak("")
	pers1 := Man{"David", "Tebet"}
	fmt.Println(pers1)
	helper.PrintBreak("")
	pers1.Married()
	fmt.Println(pers1)

	helper.PrintBreak("")
	pers2 := Man{"Eko", "Tebet"}
	fmt.Println(pers2)
	helper.PrintBreak("")
	pers2.Married()
	fmt.Println(pers2)
	fmt.Println(pers1) //pers1 tidak berubah , aman , hemat memory, func tidak perlu return
}
