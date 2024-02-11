package main

import (
	"fmt"
	"log"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	/*
		os.Stdin is what we'll use in main to capture the user's input.
		It is a *File under the hood which means it implements io.Reader
		which as we know by now is a handy way of capturing text
	*/
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
