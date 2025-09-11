package main

import (
	"fmt"
	"github.com/ingemar-fei/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading congfig : %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("fei")
	if err != nil {
		log.Fatalf("error setting username : %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading congfig : %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
}
