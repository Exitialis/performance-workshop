package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_GetPlaylistById(t *testing.T) {
	storage := NewStorage()
	album, ok := storage.GetPlaylistById(2)
	require.True(t, ok)
	require.Equal(t, Playlist{
		ID:     2,
		UserID: 19,
		Logo:   "http://placehold.it/32x32",
	}, album)
}
