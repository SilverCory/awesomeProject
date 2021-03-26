package handler_help

import (
	"fmt"
	"github.com/SilverCory/awesomeProject/command"
)

var (
	_ command.Command = new(HelpCommand)
)

type HelpCommand struct {
	cm command.Manager
}

func (hc *HelpCommand) Handle(kek interface{}) error {
	fmt.Println(hc.cm.GetAllCommands())
	return nil
}