package handlers

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type AlbumClient interface {
	GetAlbumById(id uint64) (*storage.Album, error)
}

type PlaylistClient interface {
	GetPlaylistById(id uint64) (*storage.Playlist, error)
}

type SingersClient interface {
	GetSingerById(id uint64) (*storage.Singer, error)
}

type SongsClient interface {
	GetSongById(id uint64) (*storage.Song, error)
}

type UsersClient interface {
	GetUserById(id uint64) (*storage.User, error)
}

type PlaylistSongsStorage interface {
	GetPlaylistLikes(id uint64) ([]uint64, bool)
	GetPlaylistSongsById(id uint64) ([]uint64, bool)
}
