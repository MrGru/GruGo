package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/error_handle/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
