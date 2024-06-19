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

const postAPIURL = "https://dummyapi.io/data/v1/post"

func fetchPosts(page int) ([]models.Post, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", postAPIURL, page), nil)
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
	var postResponse models.PostResponse
	if err := json.Unmarshal(body, &postResponse); err != nil {
		return nil, err
	}
	return postResponse.Data, nil
}

func ScrapePosts(workers int) {
	var wg sync.WaitGroup
	pageQueue := make(chan int)

	for workerID := 1; workerID <= workers; workerID++ {
		wg.Add(1)
		go workerScrapePosts(workerID, pageQueue, &wg)
	}

	for page := 1; page <= 10; page++ {
		pageQueue <- page
	}

	close(pageQueue)
	wg.Wait()
}

func workerScrapePosts(workerID int, pageQueue <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for page := range pageQueue {
		posts, err := fetchPosts(page)
		if err != nil {
			log.Printf("Worker %d: error fetching posts from page %d: %v", workerID, page, err)
			continue
		}

		printPosts(workerID, posts)
	}
}

func printPosts(workerID int, posts []models.Post) {
	for _, post := range posts {
		fmt.Printf("Worker %d: Posted by %s %s: %s Likes: %d Tags: %v Date posted: %s\n",
			workerID, post.User.FirstName, post.User.LastName, post.Text, post.Likes, post.Tags, post.PublishDate)
	}
}
