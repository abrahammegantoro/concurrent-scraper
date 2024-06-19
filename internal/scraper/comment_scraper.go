package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/abrahammegantoro/concurrent-scraper/internal/config"
	"github.com/abrahammegantoro/concurrent-scraper/internal/models"
)

const commentAPIURL = "https://dummyapi.io/data/v1/comment"

func fetchComments(page int) ([]models.Comment, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", commentAPIURL, page), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("app-id", config.AppID)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var commentResponse models.CommentResponse
	if err := json.Unmarshal(body, &commentResponse); err != nil {
		return nil, err
	}
	return commentResponse.Data, nil
}

func ScrapeComments(workers int) {
	var wg sync.WaitGroup
	pageQueue := make(chan int)

	for workerID := 1; workerID <= workers; workerID++ {
		wg.Add(1)
		go workerScrapeComments(workerID, pageQueue, &wg)
	}

	for page := 1; page <= 10; page++ {
		pageQueue <- page
	}

	close(pageQueue)
	wg.Wait()
}

func workerScrapeComments(workerID int, pageQueue <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for page := range pageQueue {
		comments, err := fetchComments(page)
		if err != nil {
			log.Printf("Worker %d: error fetching comments from page %d: %v", workerID, page, err)
			continue
		}

		printComments(workerID, comments)
	}
}

func printComments(workerID int, comments []models.Comment) {
	for _, comment := range comments {
		fmt.Printf("Worker %d: Commented by %s %s: %s Post ID: %s Date posted: %s \n",
			workerID, comment.Owner.FirstName, comment.Owner.LastName, comment.Message, comment.Post, comment.PublishDate)
	}
}
