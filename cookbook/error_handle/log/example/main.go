package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/error_handle/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
