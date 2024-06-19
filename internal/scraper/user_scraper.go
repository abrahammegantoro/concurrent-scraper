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

const userAPIURL = "https://dummyapi.io/data/v1/user"

func fetchUsers(page int) ([]models.UserPreview, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", userAPIURL, page), nil)
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
	var userResponse models.UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, err
	}
	return userResponse.Data, nil
}

func fetchUserDetails(userID string) (models.UserFull, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", userAPIURL, userID), nil)
	if err != nil {
		return models.UserFull{}, err
	}
	req.Header.Set("app-id", config.AppID)
	resp, err := client.Do(req)
	if err != nil {
		return models.UserFull{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.UserFull{}, err
	}
	var userResponse models.UserFull
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return models.UserFull{}, err
	}
	return userResponse, nil
}

func ScrapeUsers(workers int) {
	var wg sync.WaitGroup
	pageQueue := make(chan int)

	for workerID := 1; workerID <= workers; workerID++ {
		wg.Add(1)
		go workerScrapeUsers(workerID, pageQueue, &wg)
	}

	for page := 1; page <= 10; page++ {
		pageQueue <- page
	}

	close(pageQueue)
	wg.Wait()
}

func workerScrapeUsers(workerID int, pageQueue <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for page := range pageQueue {
		users, err := fetchUsers(page)
		if err != nil {
			log.Printf("Worker %d: error fetching users from page %d: %v", workerID, page, err)
			continue
		}

		for _, userPreview := range users {
			userDetail, err := fetchUserDetails(userPreview.ID)
			if err != nil {
				log.Printf("Worker %d: error fetching user details for user ID %s: %v", workerID, userPreview.ID, err)
				continue
			}

			printUserDetails(workerID, userDetail)
		}
	}
}

func printUserDetails(workerID int, userDetail models.UserFull) {
	fmt.Printf("Worker %d: User name %s %s %s %s %s\n",
		workerID, userDetail.Title, userDetail.FirstName, userDetail.LastName, userDetail.Email, userDetail.Gender)
}
