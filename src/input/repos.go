package input

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"codehub/src/database"
	u "codehub/src/utility"

	"github.com/google/go-github/v50/github"
)

func GetGithub(username string, db *sql.DB) {
	r := bufio.NewReader(os.Stdin)
	client := github.NewClient(nil)

	opts := &github.RepositoryListOptions{
		Type: "public",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 1000,
		},
	}

	for {
		repos, _, err := client.Repositories.List(context.Background(), username, opts)
		if err != nil {
			log.Fatal(err)
		}

		if len(repos) == 0 {
			fmt.Println("No repositories found.")
			return
		}

		for _, repo := range repos {
			fmt.Printf("%s (%s)\n", *repo.Name, *repo.HTMLURL)
			fmt.Printf("Category: ")
			Category := u.Read(r)

			fmt.Printf("Technologies (separate with space): ")
			tech := u.Read(r)

			Technologies := strings.Split(tech, " ")

			p := database.Project{ID: 1, Name: *repo.Name, Category: Category, Url: *repo.HTMLURL, Technologies: Technologies}

			CreateProject(p, db)

			for i := 0; i < 3; i++ {
				fmt.Print("\033[A")
				fmt.Print("\033[K")
			}
		}

		if len(repos) < opts.PerPage {
			break
		}

		opts.Page++
	}
}
