package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jtrotsky/wise-cli/pkg/cmd/wise"
)

func main() {
	cliRoot := filepath.Base(os.Args[0])
	err := wise.NewCommand(cliRoot).Execute()
	if err != nil {
		log.Fatal(err)
	}
}
