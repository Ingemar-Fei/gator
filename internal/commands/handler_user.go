package commands

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ingemar-fei/gator/internal/database"
	_ "github.com/lib/pq"
	"time"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:\t\t\t%v\n", user.ID)
	fmt.Printf(" * Name:\t\t%v\n", user.Name)
}

func UserRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error setting username : %v", err)
	}
	fmt.Println("User created successfully")
	printUser(user)
	return nil
}

func UserLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	name := cmd.Args[0]
	_, err := s.DB.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error setting username : %v", err)
	}
	fmt.Printf("Switched to user %s\n", name)
	return nil
}
