package handler_help_cyclic

import (
	"fmt"
	"github.com/SilverCory/awesomeProject/bot"
	"github.com/SilverCory/awesomeProject/command"
)

var (
	_ command.Command = new(HelpCommand)
)

type HelpCommand struct {
	CM *bot.Manager
}

func (hc *HelpCommand) Handle(kek interface{}) error {
	fmt.Println(hc.CM.GetAllCommands())
	return nil
}