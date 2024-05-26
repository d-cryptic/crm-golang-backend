package models

// Event represents an upcoming event
type Event struct {
    Title string `json:"title"`
    Date  string `json:"date"`
}
