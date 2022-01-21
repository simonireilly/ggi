package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/ktr0731/go-fuzzyfinder"
)

// gitignore - A struct that holds names and metadata for gitignore files
type gitignore struct {
	name    string
	content string
}

func main() {
	box := packr.NewBox("./gitignore")

	gitignores := getGitignoreFiles(box)

	idx := searchGitignores(gitignores)

	c := []byte(gitignores[idx[0]].content)

	writeGitignoreFile(c)
}

// getGitignoreFiles needs to pack the binaries for the gitignores, fetch them from the box
// then return them as a slice of []gitignore structs
func getGitignoreFiles(b packr.Box) []gitignore {
	var g []gitignore
	list := b.List()

	for _, file := range list {
		// Filter out non-gitignores
		if !strings.HasSuffix(file, ".gitignore") {
			continue
		}

		content, err := b.FindString(file)
		if err != nil {
			log.Fatal(err)
		}

		g = append(g, gitignore{name: file, content: content})
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

func writeGitignoreFile(c []byte) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(dir+"/.gitignore", c, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
