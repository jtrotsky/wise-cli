package main

import (
	"log"
	"os"

	"github.com/jtrotsky/wise-cli/pkg/cmd/wise"
)

func main() error {
	err := wise.NewCommand(os.Args[0]).Execute()
	if err != nil {
		log.Fatal(err)
	}
}
