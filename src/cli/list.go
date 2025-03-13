package cli

import (
	db "codehub/src/database"
	out "codehub/src/output"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func List(db *sql.DB) {
	proj, _ := out.GetProjects(db)

	fmt.Println()
	PrintProjects(proj)
	fmt.Println()

}

func PrintProjects(projects []db.Project) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Category", "URL", "Technologies"})

	for _, p := range projects {
		tech := strings.Join(p.Technologies, ", ")
		table.Append([]string{
			fmt.Sprintf("%-6d", p.ID),
			fmt.Sprintf("%-20s", p.Name),
			fmt.Sprintf("%-20s", p.Category),
			fmt.Sprintf("%-40s", p.Url),
			fmt.Sprintf("%-40s", tech),
		})
	}

	table.SetBorder(false)  // Brak obramowań
	table.SetRowLine(false) // Brak linii pomiędzy wierszami
	// table.SetCenterSeparator("") // Brak separatorów kolumn
	// table.SetColumnSeparator("") // Brak separatorów kolumn
	// table.SetRowSeparator("")  // Brak separatorów wierszy
	// table.SetHeaderLine(false) // Brak linii pod nagłówkiem
	table.Render()
}
