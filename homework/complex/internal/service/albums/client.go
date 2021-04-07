package albums

import (
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"math/rand"
	"time"
)

type AlbumsClient struct {
	storage AlbumsStorage
}

func New(storage AlbumsStorage) *AlbumsClient {
	rand.Seed(time.Now().Unix())
	return &AlbumsClient{
		storage: storage,
	}
}

// Метод "ходит" в сервис пользователей и подгружает пользователя
func (uc *AlbumsClient) GetAlbumById(id uint64) (*storage.Album, error) {
	user, ok := uc.storage.GetAlbumById(id)
	// Имитируем сетевой запрос в сервис
	time.Sleep(time.Millisecond * 10)

	if !ok {
		return nil, albumNotFoundError
	}

	// Имитируем ошибки сервиса пользователей с 1% шансом :-)
	val := rand.Intn(100)
	if val == 1 {
		return nil, albumServiceInternalError
	}

	return &user, nil
}
