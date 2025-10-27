package main

import (
	"database/sql"
	"github.com/ingemar-fei/gator/internal/command"
	"github.com/ingemar-fei/gator/internal/config"
	"github.com/ingemar-fei/gator/internal/database"
	"github.com/ingemar-fei/gator/internal/util"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	var err error
	coms := command.ComBook{}
	coms.Register("login", command.HandlerLogin)
	coms.Register("register", command.HandlerRegister)
	coms.Register("reset", command.ResetUsersHandler)
	coms.Register("users", command.ListUsersHandler)
	coms.Register("agg", command.RSSFetchHandler)
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	if util.DebugMode() {
		config.PrintConfig(&cfg)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("connect database failed : %v", err)
	}
	defer db.Close()
	queries := database.New(db)
	runState := &command.State{
		DBQueries: queries,
		CFG:       &cfg,
	}
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments were provided.")
	}
	comName := os.Args[1]
	comArgs := os.Args[2:]
	err = coms.Run(runState, command.Com{
		Name: comName,
		Args: comArgs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
