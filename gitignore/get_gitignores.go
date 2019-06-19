package gitignore

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
)

func GetGitignores() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(repos)
	for r := range repos {
		fmt.Println(r)
	}
}
