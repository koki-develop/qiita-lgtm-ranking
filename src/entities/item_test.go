package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * ItemsHasLGTM()
 */

func TestItems_HasLGTM(t *testing.T) {
	items := &Items{
		{Title: "1", LikesCount: 1},
		{Title: "2", LikesCount: 0},
		{Title: "3", LikesCount: 10},
		{Title: "4", LikesCount: 0},
		{Title: "5", LikesCount: 100},
	}

	assert.Equal(t, &Items{
		{Title: "1", LikesCount: 1},
		{Title: "3", LikesCount: 10},
		{Title: "5", LikesCount: 100},
	}, items.HasLGTM())
}

/*
 * Items.SortByLikesCount()
 */

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
