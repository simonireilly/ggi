package main

import (
	"github.com/gobuffalo/packr"
	"github.com/simonireilly/go-gitignore-it/internal/app"
	"github.com/simonireilly/go-gitignore-it/internal/ports"
)

func main() {
	// Define ports
	var core ports.FilesPort

	// Assign Adapters to ports
	core = packr.NewBox("../gitignore")

	// Use case
	app.Run(core)
}
