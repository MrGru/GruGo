package main

import (
	"github.com/MrGru/GruGo/cookbook/distributed_systems/orchestrate"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}
	if err := orchestrate.ConnectAndQuery(session); err != nil {
		panic(err)
	}
}
