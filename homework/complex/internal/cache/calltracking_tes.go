package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLruCache_Calltracking(t *testing.T) {
	lru := NewLruCache()
	err := lru.SetCalltracking("70000001", "79251138594")
	assert.NoError(t, err)
	err = lru.SetCalltracking("70000002", "79251138592")
	assert.NoError(t, err)

	found, notFound := lru.GetCalltracking([]RealPhone{"70000001", "70000002", "70000003"})

	expectedFound := map[RealPhone]VirtualPhone{"70000001": "79251138594", "70000002": "79251138592"}
	assert.Equal(t, expectedFound, found)

	expectedNotFound := []RealPhone{"70000003"}
	assert.Equal(t, expectedNotFound, notFound)

	t.Log(found, notFound)
}
