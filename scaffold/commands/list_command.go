package scaffold

import (
	"fmt"
	"prismarin/model"
	"prismarin/scaffold"
)

type ListScaffoldCommand struct{}

func (s ListScaffoldCommand) Execute(args []string) {
	fmt.Println("Currently available scaffolds:")
	for category, scaffolds := range scaffold.Scaffolds {
		fmt.Printf("  %s:\n", category)
		for _, scaffold := range scaffolds {
			fmt.Printf("    - %s\n", scaffold.GetName())
		}
	}
}

func (s ListScaffoldCommand) GetName() string {
	return "list"
}

func (s ListScaffoldCommand) GetChildren() []model.Command {
	return []model.Command{}
}
