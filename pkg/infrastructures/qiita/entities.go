package qiita

import (
	"sort"
	"time"
)

type Item struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	LikesCount int       `json:"likes_count"`
	URL        string    `json:"url"`
	User       User      `json:"user"`
	Tags       Tags      `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`

	StockersCount int `json:"-"`
}

type Items []*Item

func (items Items) FilterWithMinLiked(min int) Items {
	var rslt Items
	for _, item := range items {
		if item.LikesCount >= min {
			rslt = append(rslt, item)
		}
	}
	return rslt
}

func (items Items) Sort() {
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].LikesCount > items[j].LikesCount {
			return true
		}
		if items[i].LikesCount == items[j].LikesCount {
			if items[i].StockersCount > items[j].StockersCount {
				return true
			}
			if items[i].StockersCount == items[j].StockersCount {
				if items[i].CreatedAt.After(items[j].CreatedAt) {
					return true
				}
			}
		}

		return false
	})
}

type User struct {
	ID              string `json:"id"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Users []*User

type Tag struct {
	Name string `json:"name"`
}

type Tags []*Tag
