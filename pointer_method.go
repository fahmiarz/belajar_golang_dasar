package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

}
func main() {
	fahmi := Man{"Fahmi"}
	fahmi.Married()

	fmt.Println(fahmi.Name)
}
