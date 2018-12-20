package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/error_handle/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
