// cmd/nextturbo/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nextturbo/internal/nextturbo"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nextturbo.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
