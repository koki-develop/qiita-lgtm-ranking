package qiita

import "time"

type Item struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	LikesCount int       `json:"likes_count"`
	URL        string    `json:"url"`
	User       User      `json:"user"`
	Tags       Tags      `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`

	Stockers Users `json:"-"`
}

type Items []*Item

type User struct {
	ID              string `json:"id"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Users []*User

type Tag struct {
	Name string `json:"name"`
}

type Tags []*Tag
