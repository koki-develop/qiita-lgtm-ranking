package entities

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

	Stockers Users `json:"-"`
}

type Items []*Item

func (item *Item) HasLGTM() bool {
	return item.LikesCount > 0
}

func (items Items) FilterOnlyHasLGTM() Items {
	var rtn Items
	for _, item := range items {
		if item.HasLGTM() {
			rtn = append(rtn, item)
		}
	}
	return rtn
}

func (items Items) Sort() {
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].LikesCount > items[j].LikesCount {
			return true
		}
		if items[i].LikesCount == items[j].LikesCount {
			if len(items[i].Stockers) > len(items[j].Stockers) {
				return true
			}
			if len(items[i].Stockers) == len(items[j].Stockers) {
				if items[i].CreatedAt.After(items[j].CreatedAt) {
					return true
				}
			}
		}

		return false
	})
}
