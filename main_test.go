package main

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestGetGitignoreFiles(t *testing.T) {
	var folder = "./test/fixtures"
	var file = "example_2.gitignore"

	gitignores := getGitignoreFiles("./test/fixtures")
	content, err := ioutil.ReadFile(path.Join(folder, file))

	if err != nil {
		t.Errorf("Error reading file '%s", file)
	}

	if len(gitignores) != 2 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(gitignores), 2)
	}

	if gitignores[1].content != string(content) {
		t.Errorf("Content was incorrect, got: '%s', want: '%s'.", gitignores[1].content, "## Example text - 2")
	}
}
