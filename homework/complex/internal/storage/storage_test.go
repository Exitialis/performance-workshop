package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadData(t *testing.T) {
	s := Storage{
		users: make(map[uint64]User),
	}

	err := s.loadUsers()
	require.NoError(t, err)
	require.Equal(t, 20, len(s.users))
}

func TestInit(t *testing.T) {
	s := NewStorage()

	require.Equal(t, 20, len(s.users))
	require.Equal(t, 100, len(s.songs))
	require.Equal(t, 20, len(s.albums))
	require.Equal(t, 20, len(s.playlists))
	require.Equal(t, 10, len(s.singers))
	require.Equal(t, 20, len(s.userPlaylistLikes))
	require.Equal(t, 21, len(s.playlistSongs))
}
