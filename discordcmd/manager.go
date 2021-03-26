package discordcmd

import (
	"github.com/SilverCory/awesomeProject/command"
)

var (
	_ command.Manager = new(Manager)
)

type Manager struct {
	cmdMap map[string]command.Command
}

func (m *Manager) GetCommand(name string) command.Command {
	return m.cmdMap[name]
}

func (m *Manager) GetAllCommands() (ret []command.Command) {
	for _, v := range m.cmdMap {
		ret = append(ret, v)
	}

	return ret
}

