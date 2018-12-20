package main

import "github.com/MrGru/GruGo/cookbook/databases_and_storage/mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
