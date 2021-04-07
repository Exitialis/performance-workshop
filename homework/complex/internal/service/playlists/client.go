package playlists

import (
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"math/rand"
	"time"
)

type PlaylistClient struct {
	storage PlaylistStorage
}

func New(storage PlaylistStorage) *PlaylistClient {
	rand.Seed(time.Now().Unix())
	return &PlaylistClient{
		storage: storage,
	}
}

// Метод "ходит" в сервис плейлистов и подгружает плейлист
func (uc *PlaylistClient) GetPlaylistById(id uint64) (*storage.Playlist, error) {
	playlist, ok := uc.storage.GetPlaylistById(id)
	// Имитируем сетевой запрос в сервис
	time.Sleep(time.Millisecond * 30)

	if !ok {
		return nil, playlistNotFoundError
	}

	// Имитируем ошибки сервиса плейлистов с 1% шансом :-)
	val := rand.Intn(100)
	if val == 1 {
		return nil, playlistServiceInternalError
	}

	return &playlist, nil
}
