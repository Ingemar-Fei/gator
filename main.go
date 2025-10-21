package main

import (
	"github.com/ingemar-fei/gator/internal/command"
	"github.com/ingemar-fei/gator/internal/config"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	var err error
	coms := command.ComBook{}
	coms.Register("Login", command.HandlerLogin)
	runState := command.State{}
	runState.CFG, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	DebugMessage("--- read success ---\nuser: [%v]\ndburl: [%v]\n--------------------\n", runState.CFG.CurUserName, runState.CFG.DBUrl)
	if len(os.Args) < 3 {
		log.Fatal("not enough arguments were provided.")
	}
	comName := os.Args[1]
	comArgs := os.Args[2:]
	err = coms.Run(&runState, command.Com{
		Name: comName,
		Args: comArgs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
