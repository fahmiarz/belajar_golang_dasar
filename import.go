package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Fahmi")
	fmt.Println(result)

	fmt.Println(helper.Application)         //bisa
	fmt.Println(helper.sayGoodBye("Fahmi")) //tidak bisa
	fmt.Println(helper.version)             //tidak bisa
}
