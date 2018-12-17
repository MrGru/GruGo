package main

import "github.com/MrGru/GruGo/cookbook/io/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
