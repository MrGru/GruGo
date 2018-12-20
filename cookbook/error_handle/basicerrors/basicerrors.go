package basicerrors

import (
	"errors"
	"fmt"
)

var ErrorTyped = errors.New("this is a typed error")

func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)
	err = ErrorTyped
	fmt.Println("typed error: ", err)
}
