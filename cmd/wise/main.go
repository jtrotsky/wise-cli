package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jtrotsky/wise-cli/pkg/cmd"
)

func main() {
	base := filepath.Base(os.Args[0])
	err := cmd.NewCommand(base).Execute()
	if err != nil {
		log.Fatal(err)
	}
}
