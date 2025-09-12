package commands

import (
	"errors"
	"github.com/ingemar-fei/gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type CommandBook struct {
	ValidCommand map[string]func(*State, Command) error
}

func (cb *CommandBook) Register(name string, f func(*State, Command) error) {
	cb.ValidCommand[name] = f
}

func (cb *CommandBook) Run(s *State, cmd Command) error {
	f, ok := cb.ValidCommand[cmd.Name]
	if !ok {
		return errors.New("Unknown Command")
	}
	return f(s, cmd)
}
