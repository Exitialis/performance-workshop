package albums

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type AlbumsStorage interface {
	GetAlbumById(id uint64) (storage.Album, bool)
}
