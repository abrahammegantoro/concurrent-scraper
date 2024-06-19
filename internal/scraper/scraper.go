package scraper

import (
    "sync"
)

func ScrapeData(workers int) {
    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        defer wg.Done()
        ScrapeUsers(workers)
    }()

    go func() {
        defer wg.Done()
        ScrapePosts(workers)
    }()

    go func() {
        defer wg.Done()
        ScrapeComments(workers)
    }()

    wg.Wait()
}
