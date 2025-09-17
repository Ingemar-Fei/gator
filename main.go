package main

import (
	"fmt"
	"github.com/ingemar-fei/gator/internal/commands"
	"github.com/ingemar-fei/gator/internal/config"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading congfig : %v", err)
	}
	programState := &commands.State{
		Cfg: &cfg,
	}
	cmdBook := commands.CommandBook{
		ValidCommand: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmdBook.Register("login", commands.UserLogin)
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmdBook.Run(programState, commands.Command{
		Name: cmdName,
		Args: cmdArgs,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
