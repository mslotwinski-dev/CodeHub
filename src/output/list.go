package output

import (
	"database/sql"
	"encoding/json"
	"fmt"

	database "codehub/src/database"
)

func GetProjects(db *sql.DB) ([]database.Project, error) {
	rows, err := db.Query("SELECT id, name, category, url, technologies FROM projects")
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %v", err)
	}
	defer rows.Close()

	var projects []database.Project

	for rows.Next() {
		var project database.Project
		var technologiesJSON string

		if err := rows.Scan(&project.ID, &project.Name, &project.Category, &project.Url, &technologiesJSON); err != nil {
			return nil, fmt.Errorf("failed to scan project: %v", err)
		}

		err := json.Unmarshal([]byte(technologiesJSON), &project.Technologies)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal technologies: %v", err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}
