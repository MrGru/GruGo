package main

import (
	"github.com/MrGru/GruGo/cookbook/databases_and_storage/database"
	"github.com/MrGru/GruGo/cookbook/databases_and_storage/dbinterface"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	if err := dbinterface.Exec(db); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
