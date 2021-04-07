package songs

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type SongsStorage interface {
	GetSongById(id uint64) (storage.Song, bool)
}
