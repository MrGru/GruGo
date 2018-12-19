package main

import "github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
