package main

import "github.com/MrGru/GruGo/cookbook/error_handle/global"

func main() {

	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
