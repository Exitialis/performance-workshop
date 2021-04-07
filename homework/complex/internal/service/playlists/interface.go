package playlists

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type PlaylistStorage interface {
	GetPlaylistById(id uint64) (storage.Playlist, bool)
	GetPlaylistLikes(userId uint64) ([]uint64, bool)
}
