package main

import (
	"github.com/gobuffalo/packr"
	"github.com/simonireilly/go-gitignore-it/internal/app"
)

func main() {
	// Ports
	b := packr.NewBox("./gitignore")

	// Use case
	app.Run(b)
}
