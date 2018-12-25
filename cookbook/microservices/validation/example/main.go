package main

import (
	"fmt"
	"net/http"

	"github.com/MrGru/GruGo/cookbook/microservices/validation"
)

func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listen on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
