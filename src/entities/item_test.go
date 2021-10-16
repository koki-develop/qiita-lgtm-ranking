package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItem_HasLGTM(t *testing.T) {
	t.Run("return true when item has lgtm", func(t *testing.T) {
		items := Items{
			{LikesCount: 1},
			{LikesCount: 10},
			{LikesCount: 100},
		}
		for _, item := range items {
			assert.True(t, item.HasLGTM())
		}
	})
	t.Run("return false when item hasn't lgtm", func(t *testing.T) {
		items := Items{
			{LikesCount: 0},
		}
		for _, item := range items {
			assert.False(t, item.HasLGTM())
		}
	})
}

func TestItems_FilterOnlyHasLGTM(t *testing.T) {
	t.Run("return items only has lgtm", func(t *testing.T) {
		items := Items{
			{Title: "TITLE_1", LikesCount: 1},
			{Title: "TITLE_2", LikesCount: 0},
			{Title: "TITLE_3", LikesCount: 100},
			{Title: "TITLE_4", LikesCount: 1000},
		}
		assert.Equal(t, Items{
			{Title: "TITLE_1", LikesCount: 1},
			{Title: "TITLE_3", LikesCount: 100},
			{Title: "TITLE_4", LikesCount: 1000},
		}, items.FilterOnlyHasLGTM())
	})
}

func TestItems_Sort(t *testing.T) {
	items := Items{
		{Title: "1", LikesCount: 1, Stockers: Users{}},
		{Title: "2", LikesCount: 0, Stockers: Users{}},
		{Title: "3", LikesCount: 10, Stockers: Users{}},
		{Title: "4", LikesCount: 0, Stockers: Users{}},
		{Title: "5", LikesCount: 100, Stockers: Users{}},
		{Title: "6", LikesCount: 1, Stockers: Users{}},
		{Title: "7", LikesCount: 1, Stockers: Users{}},
		{Title: "8", LikesCount: 1, Stockers: Users{{}}, CreatedAt: time.Date(1998, 1, 30, 0, 0, 0, 0, time.UTC)},
		{Title: "9", LikesCount: 1, Stockers: Users{{}}, CreatedAt: time.Date(1998, 1, 31, 0, 0, 0, 0, time.UTC)},
	}

	items.Sort()

	var ttls []string
	for _, item := range items {
		ttls = append(ttls, item.Title)
	}

	assert.Equal(t, []string{
		"5",
		"3",
		"9",
		"8",
		"1",
		"6",
		"7",
		"2",
		"4",
	}, ttls)
}
