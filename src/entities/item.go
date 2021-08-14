package entities

import (
	"sort"
	"time"
)

type Item struct {
	Title      string    `json:"title"`
	LikesCount int       `json:"likes_count"`
	URL        string    `json:"url"`
	User       User      `json:"user"`
	Tags       Tags      `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
}

type Items []*Item

func (items *Items) HasLGTM() *Items {
	rtn := &Items{}

	for _, item := range *items {
		if item.LikesCount > 0 {
			*rtn = append(*rtn, item)
		}
	}

	return rtn
}

func (items *Items) SortByLikesCount() {
	sort.SliceStable(*items, func(i, j int) bool {
		return (*items)[i].LikesCount > (*items)[j].LikesCount
	})
}
