package templates

import (
	"bufio"
	"codehub/src/database"
	"codehub/src/output"
	u "codehub/src/utility"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetCategories(projects []database.Project) map[string]string {
	r := bufio.NewReader(os.Stdin)

	C := make(map[string]string)

	for _, project := range projects {
		if _, exists := C[project.Category]; !exists {
			fmt.Printf("Description for category '%s': ", project.Category)
			C[project.Category] = u.Read(r)
		}
	}

	return C
}

func Generate(db *sql.DB, params Readme) {
	projects, _ := output.GetProjects(db)

	C := GetCategories(projects)

	readme := GetHeader(params)

	readme += GetProjects(projects, C)

	readme += GetFooter()

	file, err := os.Create("_README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(readme)
	if err != nil {
		log.Fatal(err)
	}
}
