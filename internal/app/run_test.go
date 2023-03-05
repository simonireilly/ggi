package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_ports "github.com/simonireilly/go-gitignore-it/mocks/ports"
)

func TestGetGitignoreFiles(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	box := mock_ports.NewMockFilesPort(mockCtl)

	l := []string{
		"test/fixtures/example_1.gitignore",
		"test/fixtures/example_2.gitignore",
	}
	c := "contents of the file"

	box.EXPECT().List().Return(l).Times(1)
	box.EXPECT().FindString(gomock.Eq("test/fixtures/example_1.gitignore")).Return(c, nil).Times(1)
	box.EXPECT().FindString(gomock.Eq("test/fixtures/example_2.gitignore")).Return(c, nil).Times(1)

	gitignores := getGitignoreFiles(box)

	if len(gitignores) != 2 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(gitignores), 2)
	}

	if gitignores[0].content != string(c) {
		t.Errorf("Content was incorrect, got: '%s', want: '%s'.", gitignores[1].content, c)
	}
}
