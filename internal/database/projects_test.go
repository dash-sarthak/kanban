package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func projectSetup() (Author, error) {
	dummyAuthor := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "John Doe",
		Username: "john_doe",
	}
	author, err := testQueries.CreateAuthor(context.Background(), dummyAuthor)
	return author, err
}

func TestCreateProject(t *testing.T) {
	// Setup
	author, setupErr := projectSetup()
	require.NoError(t, setupErr)

	args := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      author.ID,
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	project, err := testQueries.CreateProject(context.Background(), args)

	require.NoError(t, err)
	require.NotZero(t, project)
	require.Equal(t, project.Name, args.Name)
	require.Equal(t, project.ID, args.ID)
	require.Equal(t, project.Author, args.Author)
	require.Equal(t, project.Description, args.Description)
	require.NotZero(t, project.CreatedAt)
	require.NotZero(t, project.UpdatedAt)

	// Teardown
	tdErr := authorTearDown(author.Username)
	require.NoError(t, tdErr)
}

func TestCreateProjectFalseAuthor(t *testing.T) {
	args := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      uuid.New(),
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	_, err := testQueries.CreateProject(context.Background(), args)

	require.Error(t, err)
}

func TestDeleteProject(t *testing.T) {
	// Setup
	author, setupErr := projectSetup()
	require.NoError(t, setupErr)

	args := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      author.ID,
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	project, _ := testQueries.CreateProject(context.Background(), args)
	deletedProject, err := testQueries.DeleteProject(context.Background(), project.ID)

	require.NoError(t, err)
	require.NotZero(t, deletedProject)
	require.Equal(t, deletedProject.Name, args.Name)
	require.Equal(t, deletedProject.Author, args.Author)

	// Teardown
	tdErr := authorTearDown(author.Username)
	require.NoError(t, tdErr)
}

func TestDeleteProjectFalseProject(t *testing.T) {
	deletedProject, err := testQueries.DeleteProject(context.Background(), uuid.New())

	require.Error(t, err)
	require.Zero(t, deletedProject)
}

func TestDeleteProjectCascade(t *testing.T) {
	// Setup
	author, setupErr := projectSetup()
	require.NoError(t, setupErr)

	args := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      author.ID,
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	project, _ := testQueries.CreateProject(context.Background(), args)
	_, _ = testQueries.DeleteAuthor(context.Background(), author.Username)
	deletedProject, err := testQueries.DeleteProject(context.Background(), project.ID)

	require.Error(t, err)
	require.Zero(t, deletedProject)
}

func TestGetProjects(t *testing.T) {
	// Setup
	author, setupErr := projectSetup()
	require.NoError(t, setupErr)

	projectArgs1 := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      author.ID,
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	projectArgs2 := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My Second Project",
		Author:      author.ID,
		Description: sql.NullString{String: "", Valid: false},
	}

	_, _ = testQueries.CreateProject(context.Background(), projectArgs1)
	_, _ = testQueries.CreateProject(context.Background(), projectArgs2)

	projects, err := testQueries.GetProjects(context.Background())

	require.NoError(t, err)
	require.NotZero(t, projects)
	require.Equal(t, len(projects), 2)
	require.Equal(t, projects[0].ID, projectArgs1.ID)
	require.Equal(t, projects[1].ID, projectArgs2.ID)
	require.Equal(t, projects[1].Description, projectArgs2.Description)

	// Teardown
	tdErr := errors.Join(authorTearDown(author.Username))
	require.NoError(t, tdErr)
}

func TestGetProjectsNoProjects(t *testing.T) {
	projects, err := testQueries.GetProjects(context.Background())

	require.NoError(t, err)
	require.NotZero(t, projects)
	require.Equal(t, projects, []Project{})
}

func TestGetProjectsByAuthor(t *testing.T) {
	// Setup
	author, setupErr := projectSetup()
	require.NoError(t, setupErr)

	projectArgs1 := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My First Project",
		Author:      author.ID,
		Description: sql.NullString{String: "This is a test project", Valid: true},
	}
	projectArgs2 := CreateProjectParams{
		ID:          uuid.New(),
		Name:        "My Second Project",
		Author:      author.ID,
		Description: sql.NullString{String: "", Valid: false},
	}

	_, _ = testQueries.CreateProject(context.Background(), projectArgs1)
	_, _ = testQueries.CreateProject(context.Background(), projectArgs2)

	projects, err := testQueries.GetProjectsByAuthor(context.Background(), author.ID)

	require.NoError(t, err)
	require.NotZero(t, projects)
	require.Equal(t, len(projects), 2)
	require.Equal(t, projects[0].Author, author.ID)
	require.Equal(t, projects[1].Author, author.ID)

	// Teardown
	tdErr := authorTearDown(author.Username)
	require.NoError(t, tdErr)
}
