package singers

import (
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"math/rand"
	"time"
)

type SingersClient struct {
	storage SingersStorage
}

func New(storage SingersStorage) *SingersClient {
	rand.Seed(time.Now().Unix())
	return &SingersClient{
		storage: storage,
	}
}

// Метод "ходит" в сервис пользователей и подгружает пользователя
func (uc *SingersClient) GetSingerById(id uint64) (*storage.Singer, error) {
	user, ok := uc.storage.GetSingerById(id)
	// Имитируем сетевой запрос в сервис
	time.Sleep(time.Millisecond * 10)

	if !ok {
		return nil, singerNotFoundError
	}

	// Имитируем ошибки сервиса пользователей с 1% шансом :-)
	val := rand.Intn(100)
	if val == 1 {
		return nil, singerServiceInternalError
	}

	return &user, nil
}
