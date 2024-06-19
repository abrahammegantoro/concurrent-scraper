package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "Concurrent Scraper",
    Short: "Simple Concurrent Scraper Application",
    Long:  `A simple concurrent scraper application to scrape data from dummy API concurrently.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize()
}
