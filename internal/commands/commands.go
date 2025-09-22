package commands

import (
	"fmt"
	"github.com/ingemar-fei/gator/internal/config"
	"github.com/ingemar-fei/gator/internal/database"
)

type State struct {
	DB  *database.Queries
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
		return fmt.Errorf("unknown command")
	}
	return f(s, cmd)
}
