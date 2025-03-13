package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	database "codehub/src/database"
	in "codehub/src/input"
)

func Read(reader *bufio.Reader) string {

	v, _ := reader.ReadString('\n')
	v = strings.TrimSpace(v)

	return v
}

func Add(db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	Hi()

	fmt.Printf("Project name: ")
	Name := Read(r)

	fmt.Printf("Category: ")
	Category := Read(r)

	fmt.Printf("URL to github: ")
	Url := Read(r)

	fmt.Printf("Technologies (separate with space): ")
	tech := Read(r)

	Technologies := strings.Split(tech, " ")

	p := database.Project{ID: 1, Name: Name, Category: Category, Url: Url, Technologies: Technologies}

	in.CreateProject(p, db)

}
