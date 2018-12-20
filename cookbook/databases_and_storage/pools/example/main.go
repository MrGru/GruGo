package main

import "github.com/MrGru/GruGo/cookbook/databases_and_storage/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
