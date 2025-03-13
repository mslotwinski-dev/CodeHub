package src

import (
	"codehub/src/cli"
	"codehub/src/database"
	out "codehub/src/output"
	"log"
	"os"
)

func Init() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	database.CreateTable(db)

	args := os.Args[1:]

	if len(args) == 0 {
		cli.Usage()
		return
	}

	switch args[0] {

	case "--version":
		cli.Version()
	case "help":
		cli.Help(args[1:])
	case "add":
		cli.Add(db)
	case "list":
		cli.List(db)
	default:
		out.Error("Invalid command")
		cli.Usage()

	}

}
