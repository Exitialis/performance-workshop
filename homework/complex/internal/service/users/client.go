package users

import (
	"github.com/exitialis/workshop/homework/complex/internal/storage"
	"math/rand"
	"time"
)

type UserClient struct {
	storage UserStorage
}

func New(storage UserStorage) *UserClient {
	rand.Seed(time.Now().Unix())
	return &UserClient{
		storage: storage,
	}
}

// Метод "ходит" в сервис пользователей и подгружает пользователя
func (uc *UserClient) GetUserById(id uint64) (*storage.User, error) {
	user, ok := uc.storage.GetUserById(id)
	// Имитируем сетевой запрос в сервис
	time.Sleep(time.Millisecond * 10)

	if !ok {
		return nil, userNotFoundError
	}

	// Имитируем ошибки сервиса пользователей с 1% шансом :-)
	val := rand.Intn(100)
	if val == 1 {
		return nil, userServiceInternalError
	}

	return &user, nil
}
