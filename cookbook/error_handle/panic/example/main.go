package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/error_handle/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
