package cli

import (
	"fmt"
)

func Hi() {
	fmt.Println("\u001B[36mCodeHub\u001B[0m - your personal project manager.")

	fmt.Println()
}

func Usage() {
	Hi()

	fmt.Println("Usage: ")

	fmt.Println()

	fmt.Println("\tcodehub <command> [arguments]")

	fmt.Println()

	fmt.Println("Commands:")

	fmt.Println("\tadd \t\tadd new project to database")
	fmt.Println("\tgenerate \t\tgenerate github readme")
	fmt.Println("\tlist \t\tdisplay list of projects")
	fmt.Println("\tremove \t\tremove project from database")

	fmt.Println()

	fmt.Println("Use 'codehub help <command>' for more information about that command.")

	fmt.Println()

}
