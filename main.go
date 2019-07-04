package main

import (
	"fmt"
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
)

// gitignore - A struct that holds names and metadata for gitignore files
type gitignore struct {
	Type        string
	Name        string
	LastUpdated string
}

// A placeholder for gitignore files
// This will be replaced by an at runtime reading of the file names
// with pointers to the location of the file
var gitignores = []gitignore{
	{"Global", "VisualStudioCode", "2019-05-16"},
	{"Global", "SublimeText 3", "2019-05-06"},
	{"Community - Linux", "Snap", "2019-02-16"},
	{"Community", "Red", "2018-05-16"},
	{"Root", "Node", "2019-06-16"},
}

func main() {
	idx, err := fuzzyfinder.FindMulti(
		gitignores,
		func(i int) string {
			return gitignores[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(
				"gitignore: %s (%s)\nAlbum: %s",
				gitignores[i].Name,
				gitignores[i].Type,
				gitignores[i].LastUpdated,
			)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}
