package main

import (
	"fmt"
)

func main() {
	const firstName string = "Rio"
	const lastName = "Jozef"
	const num = 1000
	//	Note :
	//		1.	Constant tidak dapat diubah valuenya, tidak seperti variable
	//		2.	Jika mencoba merubah nilai constant maka akan error
	// constant jika didefine dan tidak digunakan maka tidak akan error, tidak seperti variable

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(num)

	const (
		firstName2 string = "Rio"
		lastName2         = "Jozef"
		num2              = 2000
	)
}
