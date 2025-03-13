package cli

import (
	out "codehub/src/output"
	"fmt"
)

func Help(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}

	switch args[0] {
	case "add":
		HelpAdd()
	case "list":
		HelpList()
	case "remove":
		HelpRemove()
	default:
		out.Error("Unknown command: " + args[0])
		Usage()
	}
}

func HelpAdd() {
	fmt.Println("\u001B[36madd\u001B[0m - Add a new project to the database.")
	fmt.Println()
	fmt.Println("Usage: codehub add <project-name> <repository-url>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("\t<project-name> \tName of the project.")
	fmt.Println("\t<repository-url>\tURL of the project repository.")
}

func HelpList() {
	fmt.Println("\u001B[36mlist\u001B[0m - Display the list of projects.")
	fmt.Println()
	fmt.Println("Usage: codehub list")
	fmt.Println()
	fmt.Println("This command will show all the projects in the database.")
}

func HelpRemove() {
	fmt.Println("\u001B[36mremove\u001B[0m - Remove a project from the database.")
	fmt.Println()
	fmt.Println("Usage: codehub remove <project-name>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("\t<project-name> \tName of the project to remove.")
}
