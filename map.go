package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Fahmi"
	//person["address"] = "Jakarta"

	person := map[string]string{
		"name":    "Fahmi",
		"address": "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := map[string]string{
		"title":  "Buku Golang",
		"author": "Fahmi",
		"ups":    "Salah",
	}

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
