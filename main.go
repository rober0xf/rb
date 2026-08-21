package main

import (
	"log"
	"rb/cmd"
	"rb/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	if err := cmd.Execute(cfg); err != nil {
		log.Fatalf("error running Execute: %v", err)
	}
}
