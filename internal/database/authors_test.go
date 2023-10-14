package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func tearDown(username string) error {
	_, err := testQueries.DeleteAuthor(context.Background(), username)
	return err
}

func TestCreateAuthor(t *testing.T) {
	args := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "Sarthak Dash",
		Username: "dash_sarthak",
	}
	author, err := testQueries.CreateAuthor(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, args.ID, author.ID)
	require.Equal(t, args.Name, author.Name)
	require.Equal(t, args.Username, author.Username)

	require.NotZero(t, author.CreatedAt)
	require.NotZero(t, author.UpdatedAt)

	// Teardown
	tdErr := tearDown(args.Username)
	require.NoError(t, tdErr)
}

func TestDeleteAuthor(t *testing.T) {
	args := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "Sarthak Dash",
		Username: "dash_sarthak",
	}
	_, _ = testQueries.CreateAuthor(context.Background(), args)

	username, err := testQueries.DeleteAuthor(context.Background(), args.Username)

	require.NoError(t, err)
	require.Equal(t, args.Username, username)

	// Teardown
	if err != nil {
		tdErr := tearDown(args.Username)
		require.NoError(t, tdErr)
	}
}

func TestDeleteAuthorFalseUser(t *testing.T) {
	_, err := testQueries.DeleteAuthor(context.Background(), "UserThatDoesNotExist")

	require.Error(t, err)
}

func TestCreateAuthorDuplicateUsername(t *testing.T) {
	author1Args := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "Sarthak Dash",
		Username: "dash_sarthak",
	}

	_, _ = testQueries.CreateAuthor(context.Background(), author1Args)

	author2Args := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "Sarthak Dash",
		Username: "dash_sarthak",
	}
	_, err := testQueries.CreateAuthor(context.Background(), author2Args)
	require.Error(t, err)

	// Teardown
	tdErr := tearDown(author1Args.Username)
	require.NoError(t, tdErr)

}

func TestGetAuthors(t *testing.T) {
	args := CreateAuthorParams{
		ID:       uuid.New(),
		Name:     "Sarthak Dash",
		Username: "dash_sarthak",
	}
	_, _ = testQueries.CreateAuthor(context.Background(), args)

	authors, err := testQueries.GetAuthors(context.Background())

	require.NoError(t, err)
	require.NotZero(t, authors)
	require.Equal(t, 1, len(authors))
	require.Equal(t, "Sarthak Dash", authors[0].Name)
	require.Equal(t, "dash_sarthak", authors[0].Username)

	// Teardown
	tdErr := tearDown(args.Username)
	require.NoError(t, tdErr)
}
