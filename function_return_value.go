package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Fahmi Arzalega")
	fmt.Println(result)

	fmt.Println(getHello("Anjas"))
	fmt.Println(getHello("Harun"))
}
