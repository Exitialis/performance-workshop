package songs

import (
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"math/rand"
	"time"
)

type SongsClient struct {
	storage SongsStorage
}

func New(storage SongsStorage) *SongsClient {
	rand.Seed(time.Now().Unix())
	return &SongsClient{
		storage: storage,
	}
}

// Метод "ходит" в сервис пользователей и подгружает пользователя
func (uc *SongsClient) GetSongById(id uint64) (*storage.Song, error) {
	user, ok := uc.storage.GetSongById(id)
	// Имитируем сетевой запрос в сервис
	time.Sleep(time.Millisecond * 8)

	if !ok {
		return nil, songNotFoundError
	}

	// Имитируем ошибки сервиса пользователей с 1% шансом :-)
	val := rand.Intn(100)
	if val == 1 {
		return nil, songServiceInternalError
	}

	return &user, nil
}
