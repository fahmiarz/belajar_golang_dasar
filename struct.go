package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func (customer Customer) sayAddres(name string) {
	fmt.Println("Saya", name, "tinggal di", customer.Address, customer.Name)
}
func main() {
	fahmi := Customer{}
	fahmi.Name = "Fahmi Arzalega"
	fahmi.Address = "Jakarta"
	fahmi.Age = 25
	fmt.Println(fahmi)

	udin := Customer{
		Name:    "Udin",
		Address: "Jakarta",
		Age:     24,
	}
	fmt.Println(udin)

	budi := Customer{"Budi", "Jakarta", 29}
	fmt.Println(budi)

	fahmi.sayHello("Agus")
	udin.sayHello("Agus")
	budi.sayHello("Agus")

	fahmi.sayAddres("Agus")
	udin.sayAddres("Agus")
	budi.sayAddres("Agus")
}
