package helper

import "fmt"

func PrintBreak(title string) {
	switch title != "" {
	case true:
		fmt.Println("===============", title, "===============")
	default:
		fmt.Println("=========================================")
	}
}
