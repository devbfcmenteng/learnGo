package main

import (
	"bfcTest/packages/helper"
	"errors"
	"fmt"
)

func pembagian(val int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh dengan Nol")
	} else {
		res := val / pembagi
		return res, nil
	}

}

func main() {
	/*
		Golang memiliki interface yang digunakan sebagai kontrak untuk MEMBUAT error, yaitu interface error
	*/

	helper.PrintBreak("Cth Error 1")
	var contohError error = errors.New("Ups Ada Error")
	fmt.Println(contohError)

	helper.PrintBreak("Cth Error 2")
	resPembagian, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", resPembagian)
	} else {
		fmt.Println("Error =", err.Error())
	}

}
