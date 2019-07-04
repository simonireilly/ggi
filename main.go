package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ktr0731/go-fuzzyfinder"
)

// gitignore - A struct that holds names and metadata for gitignore files
type gitignore struct {
	name    string
	content string
}

func main() {
	gitignores := getGitignoreFiles("./test/fixtures")

	idx := searchGitignores(gitignores)

	c := []byte(gitignores[idx[0]].content)

	err := ioutil.WriteFile("./test/.gitignore", c, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getGitignoreFiles(folder string) []gitignore {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var g []gitignore

	for _, file := range files {
		content, err := ioutil.ReadFile(path.Join(folder, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		g = append(g, gitignore{name: file.Name(), content: string(content)})
	}
	return g
}

func searchGitignores(g []gitignore) []int {
	idx, err := fuzzyfinder.FindMulti(
		g,
		func(i int) string {
			return g[i].name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(
				"gitignore: %s \n---\n%s\n---",
				g[i].name,
				g[i].content,
			)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return idx
}
