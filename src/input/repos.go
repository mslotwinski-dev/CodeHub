package input

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"codehub/src/database"
	"codehub/src/output"
	u "codehub/src/utility"
)

func GetGithub(username string, db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	repos := output.GetGithub(username)

	for _, repo := range repos {
		fmt.Printf("%s (%s)\n", *repo.Name, *repo.HTMLURL)
		fmt.Printf("Category: ")
		Category := u.Read(r)

		fmt.Printf("Technologies (separate with space): ")

		if repo.Language != nil {
			fmt.Printf("%s ", strings.ReplaceAll(*repo.Language, " ", ""))
		}

		tech := u.Read(r)

		if repo.Language != nil {
			tech += fmt.Sprintf("%s ", strings.ReplaceAll(*repo.Language, " ", ""))
		}

		Technologies := strings.Split(tech, " ")

		p := database.Project{ID: 1, Name: *repo.Name, Category: Category, Url: *repo.HTMLURL, Technologies: Technologies}

		CreateProject(p, db)

		for i := 0; i < 3; i++ {
			fmt.Print("\033[A")
			fmt.Print("\033[K")
		}
	}
}
