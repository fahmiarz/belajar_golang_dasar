package main

import "fmt"

func endApp() {
	fmt.Println("end App")
	message := recover()
	fmt.Println("Terjadi Panic nih", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aduh Error")
	}
}
func main() {
	//runApp(false)
	runApp(true)
	fmt.Println("Fahmi Arzalega")

}
