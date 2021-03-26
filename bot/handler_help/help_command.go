package handler_help

import (
	"fmt"
	"github.com/SilverCory/awesomeProject/command"
)

var (
	_ command.Command = new(HelpCommand)
)

type HelpCommand struct {
	CM command.Manager
}

func (hc *HelpCommand) Handle(kek interface{}) error {
	fmt.Println(hc.CM.GetAllCommands())
	return nil
}