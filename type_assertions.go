package main

import "fmt"

func random() any {
	return "mantap bosku"
}
func main() {
	var result = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	//var resultInt int = result.(int) //panic
	//fmt.Println(resultInt)
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}
