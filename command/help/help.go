package help

import (
	"fmt"
	"prismarin/model"
)

type HelpCommand struct{}

func (s HelpCommand) Execute(args []string) {
	SendDefaultHelpMessage()
}

func (s HelpCommand) GetName() string {
	return "help"
}

func (s HelpCommand) GetChildren() []model.Command {
	return []model.Command{}
}

func SendDefaultHelpMessage() {
	fmt.Println("")
	fmt.Println("Prismarin:")
	fmt.Println("  Commands:")
	fmt.Println("    help:                            | View commands")
	fmt.Println("    scaffold:")
	fmt.Println("      list:                          | List all available scaffolds")
	fmt.Println("      get: <category> <name>         | Retrieve a scaffold")
	fmt.Println("")
}
