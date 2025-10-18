package command

import (
	"fmt"
	"github.com/ingemar-fei/gator/internal/config"
)

type State struct {
	CFG config.Config
}

type Com struct {
	Name string
	Args []string
}

type ComBook struct {
	comBook map[string]func(*State, Com) error
}

func (c *ComBook) Register(name string, f func(*State, Com) error) {
	if c.comBook == nil {
		c.comBook = make(map[string]func(*State, Com) error)
	}
	c.comBook[name] = f
}

func (c *ComBook) Run(s *State, com Com) error {
	fName, ok := c.comBook[com.Name]
	if !ok {
		return fmt.Errorf("error, %v is not a valid command", com.Name)
	}
	return fName(s, com)
}

func HandlerLogin(s *State, cmd Com) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("error: missing <username> for login")
	}
	s.CFG.SetUser(cmd.Args[0])
	fmt.Printf("welcome, %v\n", s.CFG.CurUserName)
	return nil
}
