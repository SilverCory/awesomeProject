package command

type Manager interface {
	GetCommand(name string) Command
	GetAllCommands() []Command
}
