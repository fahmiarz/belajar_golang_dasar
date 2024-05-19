package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You Are Bloked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Fahmi", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
