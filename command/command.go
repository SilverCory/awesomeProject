package command

type Command interface {
	Handle(discordArgsHere interface{}) error
}
