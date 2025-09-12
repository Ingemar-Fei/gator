package commands

import (
	"fmt"
)

func UserLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <username>", cmd.Name)
	}
	name := cmd.Args[0]
	err := s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error setting username : %v", err)
	}
	fmt.Printf("Switched to user %s\n", name)
	return nil
}
