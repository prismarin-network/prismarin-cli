package model

type Command interface {
	Execute(args []string)
	GetName() string
	GetChildren() []Command
}
