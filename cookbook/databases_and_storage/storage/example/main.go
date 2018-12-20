package main

import "github.com/MrGru/GruGo/cookbook/databases_and_storage/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
