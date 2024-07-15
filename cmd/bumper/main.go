package main

import (
	"github.com/go-shortcut/json-bumper/pkg/jsonhelper"
	"log"
	"os"
)

const (
	banner = `
how to use...
`
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal(banner)
	}
	jsonhelper.Cli(os.Args[1:]...)
}
