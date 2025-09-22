package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ingemar-fei/gator/internal/commands"
	"github.com/ingemar-fei/gator/internal/config"
	"github.com/ingemar-fei/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading congfig : %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening db : %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	programState := &commands.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}
	cmdBook := commands.CommandBook{
		ValidCommand: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmdBook.Register("register", commands.UserRegister)
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
