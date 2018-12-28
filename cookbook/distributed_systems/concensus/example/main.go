package main

import (
	"net/http"

	"github.com/MrGru/GruGo/cookbook/distributed_systems/concensus"
)

func main() {
	concensus.Config(3)
	http.HandleFunc("/", concensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
