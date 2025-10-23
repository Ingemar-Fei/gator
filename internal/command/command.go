package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ingemar-fei/gator/internal/config"
	"github.com/ingemar-fei/gator/internal/database"
	"github.com/ingemar-fei/gator/internal/util"
)

type State struct {
	DBQueries *database.Queries
	CFG       *config.Config
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

func checkArgsNumber(cmd Com) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("error: missing <args> for %v", cmd.Name)
	}
	return nil
}

func printUser(u database.User) {
	fmt.Printf("--- User  Info ---\n # ID : %v\n # CT : %v\n # UT : %v\n # Name: %v\n--------------------\n", u.ID, u.CreatedAt, u.UpdatedAt, u.Name)
}

func HandlerLogin(s *State, cmd Com) error {
	err := checkArgsNumber(cmd)
	if err != nil {
		return err
	}
	s.CFG.SetUser(cmd.Args[0])
	fmt.Printf("welcome, %v\n", s.CFG.CurUserName)
	return nil
}

func HandlerRegister(s *State, cmd Com) error {
	err := checkArgsNumber(cmd)
	if err != nil {
		return err
	}
	var name string
	for _, arg := range cmd.Args {
		if arg[0] != '-' {
			name = arg
			break
		}
	}
	regParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}
	user, err := s.DBQueries.CreateUser(context.Background(), regParam)
	if err != nil {
		return fmt.Errorf("create user failed: %v", err)
	}
	if util.DebugMode() {
		printUser(user)
	}
	err = s.CFG.SetUser(user.Name)
	if err != nil {
		return err
	}
	if util.DebugMode() {
		config.PrintConfig(s.CFG)
	}
	return nil
}
