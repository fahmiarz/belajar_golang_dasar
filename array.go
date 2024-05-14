package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Fahmi"
	name[1] = "Arzalega"
	name[2] = "Seazzz"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}
	fmt.Println(values)
	fmt.Println(len(values))

}
