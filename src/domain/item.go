package domain

import (
	"sort"
	"time"
)

// Item .
type Item struct {
	Title      string    `json:"title"`
	LikesCount int       `json:"likes_count"`
	URL        string    `json:"url"`
	User       User      `json:"user"`
	Tags       Tags      `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
}

// Items .
type Items []*Item

// Tag .
type Tag struct {
	Name string `json:"name"`
}

// Tags .
type Tags []*Tag

// User .
type User struct {
	ID              string `json:"id"`
	ProfileImageURL string `json:"profile_image_url"`
}

// HasLGTM .
func (items *Items) HasLGTM() *Items {
	rtn := &Items{}

	for _, item := range *items {
		if item.LikesCount > 0 {
			*rtn = append(*rtn, item)
		}
	}

	return rtn
}

// SortByLikesCount .
func (items *Items) SortByLikesCount() {
	sort.SliceStable(*items, func(i, j int) bool {
		return (*items)[i].LikesCount > (*items)[j].LikesCount
	})
}
