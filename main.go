package main

import (
	"fmt"
	"github.com/ingemar-fei/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user got: %v\n", cfg.CurUserName)
	cfg.SetUser("gator")
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user got: %v\n", cfg.CurUserName)
}
