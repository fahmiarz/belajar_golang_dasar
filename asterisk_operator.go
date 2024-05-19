package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"} //address
	address2 := &address1                                    //*address (pointer)
	address2.City = "Bandung"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2) //berubah jadi bandung

	//address2 = &Address{"Jakarta", "DK Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DK Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
