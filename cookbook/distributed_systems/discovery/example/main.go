package main

import "github.com/MrGru/GruGo/cookbook/distributed_systems/discovery"

func main() {
	if err := discovery.Exec(); err != nil {
		panic(err)
	}
}
