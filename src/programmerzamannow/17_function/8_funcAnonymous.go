package main

import (
	"fmt"
	"strings"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		nameLower := strings.ToLower(name)
		return nameLower != "admin"
	}

	registerUser("Rioio", blacklist)

	registerUser("Admin", blacklist)

	registerUser("SuperAdmin", func(name string) bool {
		nameLower := strings.ToLower(name)
		return nameLower != "superadmin"
	})

	registerUser("rio", func(name string) bool {
		nameLower := strings.ToLower(name)
		return nameLower != "superadmin"
	})

}
