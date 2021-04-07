package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_GetAlbumById(t *testing.T) {
	storage := NewStorage()
	album, ok := storage.GetAlbumById(1)
	require.True(t, ok)
	require.Equal(t, Album{
		ID:   1,
		Name: "adipisicing sit amet",
		Logo: "http://placehold.it/32x32",
	}, album)
}
