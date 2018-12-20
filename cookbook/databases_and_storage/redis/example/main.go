package main

import "github.com/MrGru/GruGo/cookbook/databases_and_storage/redis"

func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}

	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
