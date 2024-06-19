package cmd

import (
	"fmt"
	// "sync"
	"time"

	"github.com/abrahammegantoro/concurrent-scraper/internal/scraper"

	"github.com/spf13/cobra"
)

var workers int

var workerCmd = &cobra.Command{
    Use:   "worker",
    Short: "Run the scraper workers",
    Long:  `Run the scraper workers to scrape data concurrently from dummy API`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Starting %d workers...\n", workers)

        start := time.Now()
        
        scraper.ScrapeData(workers)
        
        duration := time.Since(start)
        fmt.Printf("All workers have completed their tasks. Total time taken: %s\n", duration)
    },
}

func init() {
    rootCmd.AddCommand(workerCmd)
    workerCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
}
