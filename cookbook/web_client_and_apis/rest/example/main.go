package main

import "github.com/MrGru/GruGo/cookbook/web_client_and_apis/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
