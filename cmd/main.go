package main

import (
	"github.com/simonireilly/go-gitignore-it/internal/adapters/core/files"
	"github.com/simonireilly/go-gitignore-it/internal/app"
	"github.com/simonireilly/go-gitignore-it/internal/ports"
)

func main() {
	// Define ports
	var core ports.FilesPort

	// Assign Adapters to ports
	core = files.NewAdapter()

	// Use case
	app.Run(core)
}
