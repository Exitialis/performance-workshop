package users

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type UserStorage interface {
	GetUserById(id uint64) (storage.User, bool)
}
