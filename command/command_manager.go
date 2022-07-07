package command

import (
	"prismarin/command/help"
	"prismarin/model"
	commands2 "prismarin/scaffold/commands"
	"strings"
)

var commands []model.Command

func loadCommands() {
	commands = append(commands, commands2.ScaffoldCommand{}, help.HelpCommand{})
}

func Init(args []string) {
	loadCommands()

	if len(args) >= 1 {
		index := 0
		cmd := strings.ToLower(args[index])
		for _, command := range commands {
			if strings.ToLower(command.GetName()) == cmd {
				ExecuteCommand(&index, command, args)
				return
			}
		}
		return
	}
	help.SendDefaultHelpMessage()

}

func ExecuteCommand(index *int, command model.Command, args []string) {
	*index += 1
	children := command.GetChildren()
	if len(args) >= (*index)+1 {
		message := args[*index]
		for _, cmd := range children {
			if strings.ToLower(cmd.GetName()) == message {
				ExecuteCommand(index, cmd, args)
				return
			}
		}
	}
	command.Execute(args)
}
