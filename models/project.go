package models

import (
	"github.com/dash-sarthak/kanban/internal/database"
	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Author      uuid.UUID `json:"author"`
	Description string    `json:"description"`
}

func MarshalProject(project database.Project) Project {
	return Project{
		ID:          project.ID,
		Name:        project.Name,
		Author:      project.Author,
		Description: project.Description.String,
	}
}
