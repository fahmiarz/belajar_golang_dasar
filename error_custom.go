package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}
func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}
	if id != "Fahmi" {
		return &notFoundError{"data not found"}
	}

	//ok
	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("data not found:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		//sukses
		fmt.Println("sukses")
	}
}
