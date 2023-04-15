//go:build ignore
// +build ignore

// This package takes a filesystem with symlinks and converts the symbolic files
// into real files. It exists because go embed won't embed symbolic references
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	p, err := filepath.Abs("./gitignore")

	if err != nil {
		log.Fatalf("failed to create absolute path files: %v", err)
	}

	err = filepath.Walk(p, convert_sym_to_regular)

	if err != nil {
		log.Fatalf("failed to walk files: %v", err)
	}
}

func convert_sym_to_regular(path string, info fs.FileInfo, err error) error {
	if info.Mode().Type() == os.ModeSymlink {
		fmt.Printf("Found symlink as '%s'\n", path)
		// Follow link
		l, err := filepath.EvalSymlinks(path)
		if err != nil {
			log.Fatalf("Failed to find linked file '%s': %v", info.Name(), err)
		}

		// Read link contents
		c, err := os.ReadFile(l)

		if err != nil {
			log.Fatalf("Failed to read linked file '%s': %v", l, err)
		}

		// Remove original file
		err = os.Remove(path)

		if err != nil {
			log.Fatalf("Failed to remove file '%s': %v", path, err)
		}

		// Create new file
		f, err := os.Create(path)

		if err != nil {
			log.Fatalf("Failed to create new file '%s': %v", path, err)
		}

		_, err = f.Write(c)

		if err != nil {
			log.Fatalf("Failed to write new file '%s': %v", path, err)
		}
	}

	return nil
}
