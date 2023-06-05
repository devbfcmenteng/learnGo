package main

import (
	"bfcTest/packages/helper"
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Rio",
		"address": "Tebet",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	/* Function Map
	len(map) 					-> Untuk mendapatkan jumlah data di map
	map[key] 					-> Mengambil data di map dengan key
	map[key] = value			-> mengubah data di map dengan key
	make(map[TypeKey]TypeValue) -> membuat map baru
	delete(map, key)			-> menghapus data di map dengan key
	*/
	fmt.Println("len map =", len(person))
	person["name"] = "Rio Jozef"
	person["test"] = "test"
	delete(person, "address")
	fmt.Println(person)

	helper.PrintBreak("")

	var book map[string]string = make(map[string]string)
	fmt.Println(book)
	var book2 = map[string]string{}
	fmt.Println(book2)

	test := map[string]int{
		"bil1": 12,
		"bil2": 18,
	}

	fmt.Println(test)
	fmt.Println(test["bil1"])
	test2 := map[int]string{
		1: "Rio",
		2: "Jozef",
	}
	fmt.Println(test2)
	fmt.Println(test2[1])
}
