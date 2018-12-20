package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/error_handle/structured"
)

func main() {
	fmt.Println("Logrus:")
	structured.Logrus()
	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}
