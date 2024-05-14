package main

import (
	"fmt"
)

func main() {
	name := "Fahmi"

	switch name {
	case "Fahmi":
		fmt.Println("Hello Fahmi")

	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan gak?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Telalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	name = "Arzalega"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
