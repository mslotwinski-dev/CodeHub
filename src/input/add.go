package input

import (
	db "codehub/src/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func CreateProject(project db.Project, db *sql.DB) {
	err := AddProject(project, db)
	if err != nil {
		log.Fatal(err)
	}

}

func AddProject(project db.Project, db *sql.DB) error {
	technologiesJSON, err := json.Marshal(project.Technologies)
	if err != nil {
		return fmt.Errorf("failed to marshal technologies: %v", err)
	}

	insertSQL := `INSERT INTO projects (name, category, url, technologies)
	VALUES (?, ?, ?, ?)`

	_, err = db.Exec(insertSQL,
		project.Name,
		project.Category,
		project.Url,
		technologiesJSON,
	)

	if err != nil {
		return fmt.Errorf("failed to add project: %v", err)
	}

	return nil
}
