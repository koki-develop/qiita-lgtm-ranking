package entities

import (
	"testing"

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

func TestItems_SortByLikesCount(t *testing.T) {
	items := &Items{
		{Title: "1", LikesCount: 1},
		{Title: "2", LikesCount: 0},
		{Title: "3", LikesCount: 10},
		{Title: "4", LikesCount: 0},
		{Title: "5", LikesCount: 100},
	}

	items.SortByLikesCount()

	assert.Equal(t, &Items{
		{Title: "5", LikesCount: 100},
		{Title: "3", LikesCount: 10},
		{Title: "1", LikesCount: 1},
		{Title: "2", LikesCount: 0},
		{Title: "4", LikesCount: 0},
	}, items)
}
