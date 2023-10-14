// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: projects.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (id, name, author, description)
VALUES ($1, $2, $3, $4)
RETURNING id, name, author, description, created_at, updated_at
`

type CreateProjectParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Author      uuid.UUID      `json:"author"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject,
		arg.ID,
		arg.Name,
		arg.Author,
		arg.Description,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProjects = `-- name: GetProjects :many
SELECT id, name, author, description, created_at, updated_at FROM projects
`

func (q *Queries) GetProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Author,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsByAuthor = `-- name: GetProjectsByAuthor :many
SELECT id, name, author, description, created_at, updated_at FROM projects
WHERE author = $1
`

func (q *Queries) GetProjectsByAuthor(ctx context.Context, author uuid.UUID) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjectsByAuthor, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Author,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
