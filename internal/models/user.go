package models

type UserPreview struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

type UserFull struct {
    ID           string   `json:"id"`
	Title        string   `json:"title"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Gender       string   `json:"gender"`
	Email        string   `json:"email"`
	DateOfBirth  string   `json:"dateOfBirth"`
	RegisterDate string   `json:"registerDate"`
	Phone        string   `json:"phone"`
	Picture      string   `json:"picture"`
	Location     Location `json:"location"`
}

type Location struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

type UserResponse struct {
    Data []UserPreview `json:"data"`
    Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
