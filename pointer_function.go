package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address Address = Address{"Jakarta", "DKI Jakarta", ""}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}
