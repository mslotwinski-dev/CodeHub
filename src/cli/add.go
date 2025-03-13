package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	database "codehub/src/database"
	in "codehub/src/input"
	u "codehub/src/utility"
)

func Add(db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	Hi()

	fmt.Printf("Project name: ")
	Name := u.Read(r)

	fmt.Printf("Category: ")
	Category := u.Read(r)

	fmt.Printf("URL to github: ")
	Url := u.Read(r)

	fmt.Printf("Technologies (separate with space): ")
	tech := u.Read(r)

	Technologies := strings.Split(tech, " ")

	p := database.Project{ID: 1, Name: Name, Category: Category, Url: Url, Technologies: Technologies}

	in.CreateProject(p, db)

}

func GetGithub(db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	Hi()

	fmt.Printf("Github URL: https://github.com/")
	username := u.Read(r)

	fmt.Println()

	in.GetGithub(username, db)
}
