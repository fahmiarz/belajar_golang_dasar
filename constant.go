package main

import "fmt"

func main() {
	const (
		firstName string = "Fahmi"
		lastName         = "Arzalega"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	//error
	// firstName = "asu"
	// lastName = "raimu"

}
