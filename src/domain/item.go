package domain

// Item .
type Item struct {
	Title string `json:"title"`
}

// Items .
type Items []*Item
