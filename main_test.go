package main

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/gobuffalo/packr"
)

func TestGetGitignoreFiles(t *testing.T) {
	folderName := "./test/fixtures"
	box := packr.NewBox(folderName)
	fileName := "example_2.gitignore"

	gitignores := getGitignoreFiles(box)
	content, err := ioutil.ReadFile(path.Join(folderName, fileName))

	if err != nil {
		t.Errorf("Error reading fileName '%s", fileName)
	}

	if len(gitignores) != 2 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(gitignores), 2)
	}

	if gitignores[1].content != string(content) {
		t.Errorf("Content was incorrect, got: '%s', want: '%s'.", gitignores[1].content, "## Example text - 2")
	}
}
