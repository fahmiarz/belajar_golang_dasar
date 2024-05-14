package main

import "fmt"

func main() {

	type NoKTP string

	var ktpFahmi NoKTP = "111111111"

	var contoh string = "22222222"
	var contohktp NoKTP = NoKTP(contoh)

	fmt.Println(ktpFahmi)
	fmt.Println(contohktp)
}
