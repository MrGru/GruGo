package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/math"
)

func main() {
	math.Examples()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
