package api

import (
	"fmt"
	"github.com/k0kubun/pp/v3"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func FetchUserRepos(user string) {
	//// Create a new REST client with default options

	path := fmt.Sprintf("users/%s/repos", user)

	var s []interface{}
	err := GithubClientGet(path, &s)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	pp.Println("Fetching repos", s)
	//
	//// Use the client to make authenticated requests to the GitHub API
	//// For example, list repositories for the authenticated user
	//repos, err := client.ListRepositories()
	//if err != nil {
	//	log.Fatalf("Failed to list repositories: %v", err)
	//}
	//
	//// Print the repositories
	//for _, repo := range repos {
	//	fmt.Printf("Repository: %s\n", repo.Name)
	//}
}

func TestUser(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	FetchUserRepos("ax-sh")
}
