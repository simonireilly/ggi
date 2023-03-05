// files implements ports.FilesPort using an embed file system wth go embed
package files

import (
	"embed"
	"io/fs"
	"log"
	"path/filepath"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

//go:generate cp -r ../../../../gitignore ./gitignore
//go:generate go run ../../../../tools/unpack_symlinks.go
//go:embed gitignore/**/*.gitignore
var efs embed.FS

func (a Adapter) List() (out []string) {
	root := "."

	s, err := getAllFileNames(efs, root)

	if err != nil {
		log.Fatalf("Failed to open embedded gitignore files as '%s'\n: %v", root, err)
	}

	return s
}

func (a Adapter) FindString(s string) (string, error) {
	f, err := fs.ReadFile(efs, s)
	if err != nil {
		log.Fatalf("Failed to open file '%s'\n: %v", s, err)
	}

	return string(f), nil
}

func getAllFileNames(efs embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}

	entries, err := efs.ReadDir(path)

	if err != nil {
		log.Fatalf("Failed to open embedded gitignore files as '%s'\n: %v", path, err)
	}

	for _, entry := range entries {
		// Build path
		fp := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			res, err := getAllFileNames(efs, fp)
			if err != nil {
				log.Fatalf("Failed to open embedded gitignore files as '%s'\n: %v", fp, err)
			}
			out = append(out, res...)
			continue
		}

		out = append(out, fp)
	}

	return out, err
}
