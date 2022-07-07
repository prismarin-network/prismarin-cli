package main

import (
	"os"
	"prismarin/command"
	"prismarin/scaffold"
)

func main() {
	scaffold.Init()

	args := os.Args[1:]
	command.Init(args)
}
