package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s := "Duẩn"
	fmt.Println("Number: %d", len(s))
	fmt.Print("count:", utf8.RuneCountInString(s))
}
