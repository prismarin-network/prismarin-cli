package scaffold

import (
	"fmt"
	"prismarin/model"
)

type ScaffoldCommand struct{}

func (s ScaffoldCommand) Execute(args []string) {
	fmt.Println("Yeet")
}

func (s ScaffoldCommand) GetName() string {
	return "scaffold"
}

func (s ScaffoldCommand) GetChildren() []model.Command {
	return []model.Command{ListScaffoldCommand{}, GetScaffoldCommand{}}
}
