package models

//Article defines the schema of an article object
type Article struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	FeedID      string `json:"feedID"`
}

//Articles is a collection of articles
type Articles []Article
