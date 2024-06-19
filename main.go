package main

import (
	"github.com/abrahammegantoro/concurrent-scraper/cmd"
	"github.com/abrahammegantoro/concurrent-scraper/internal/config"
)

func main() {
	config.LoadConfig()
    cmd.Execute()
}