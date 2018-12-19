package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}
