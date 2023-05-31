package main

import (
	"fmt"
	"strings"
)

func sayHelloWithFilter(name string, filter func(string) string) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

func spamFilter(name string) string {
	namefilter := strings.ToLower(name)
	if namefilter == "anjing" {
		return "..."
	} else {
		return namefilter
	}
}

func main() {
	fmt.Println(sayHelloWithFilter("Rio", spamFilter))
	fmt.Println(sayHelloWithFilter("anjing", spamFilter))
	fmt.Println(sayHelloWithFilter("Anjing", spamFilter))
	fmt.Println(sayHelloWithFilter("AnjIng", spamFilter))
}
