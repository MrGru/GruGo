package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
