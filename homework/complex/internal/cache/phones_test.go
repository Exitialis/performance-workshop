package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLruCache_Phones(t *testing.T) {
	lru := NewLruCache()
	err := lru.SetPhone(1, "79251138594")
	assert.NoError(t, err)
	err = lru.SetPhone(2, "79251138592")
	assert.NoError(t, err)

	found, notFound := lru.GetPhone([]int64{1, 2, 3})

	expectedFound := map[int64]RealPhone{1: "79251138594", 2: "79251138592"}
	assert.Equal(t, expectedFound, found)

	expectedNotFound := []int64{3}
	assert.Equal(t, expectedNotFound, notFound)

	t.Log(found, notFound)
}
