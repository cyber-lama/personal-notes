package helpers

import "fmt"

func OutputFields(s interface{}) {
	fmt.Printf("Golang struct: %#v\n", s)
}
