package scaffold

import (
	"fmt"
	"prismarin/model"
	"prismarin/scaffold"
)

type GetScaffoldCommand struct{}

func (s GetScaffoldCommand) Execute(args []string) {
	if len(args) >= 4 {
		category := args[2]
		if !scaffold.ExistCategory(category) {
			fmt.Println("This category does not exists")
			return
		}
		scaffoldName := args[3]
		scaffolds := scaffold.GetCategory(category)
		if !scaffold.ExistScaffold(scaffolds, scaffoldName) {
			fmt.Println("This scaffold does not exists")
			return
		}
		scaffold := scaffold.GetScaffold(scaffolds, scaffoldName)
		fmt.Printf("Retrieving %s from %s category\n", scaffold.GetName(), category)
		scaffold.Download()
		return
	}
	fmt.Println("Missing arguments: category and name")
}

func (s GetScaffoldCommand) GetName() string {
	return "get"
}

func (s GetScaffoldCommand) GetChildren() []model.Command {
	return []model.Command{}
}
