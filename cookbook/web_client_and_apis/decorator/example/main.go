package main

import "github.com/MrGru/GruGo/cookbook/web_client_and_apis/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
