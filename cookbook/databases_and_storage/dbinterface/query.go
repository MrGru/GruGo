package dbinterface

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/databases_and_storage/database"
)

func Query(db DB) error {
	name := "Gru"
	rows, err := db.Query("SELECT name, created FROM example WHERE name=?", name)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var e database.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}
