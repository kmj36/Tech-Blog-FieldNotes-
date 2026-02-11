package main

import (
	"log"

	"github.com/kmj36/fieldnotes-tech-blog/internal/app"
)

func main() {
	backendApp := app.New()

	if err := backendApp.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
