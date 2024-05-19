package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fahmi"
	middleName = "Seazz"
	lastName = "Arzalega"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
