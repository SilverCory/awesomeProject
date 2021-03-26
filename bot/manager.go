package bot

import (
	"github.com/SilverCory/awesomeProject/bot/handler_help"
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

func (m *Manager) RegisterDefaultCommands() {
	m.cmdMap["help"] = &handler_help.HelpCommand{CM: m}
	// m.cmdMap["help_cyclic"] = &handler_help_cyclic.HelpCommand{CM: m} // TODO uncomment for cyclic memes
}