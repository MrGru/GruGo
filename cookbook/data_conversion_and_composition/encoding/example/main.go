package main

import "github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/encoding"

func main() {
	if err := encoding.Base64Example(); err != nil {
		panic(err)
	}
	if err := encoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}
	if err := encoding.GobExample(); err != nil {
		panic(err)
	}
}
