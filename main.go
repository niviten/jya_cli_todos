package main

import (
	"doem/cmd"
	"doem/internal/db"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println("Error while init db", err)
		return
	}

	cmd.Execute()

	err = db.Close()
	if err != nil {
		fmt.Println("Error while closing db", err)
	}
}
