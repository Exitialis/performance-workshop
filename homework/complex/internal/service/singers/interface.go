package singers

import "github.com/exitialis/workshop/homework/complex/internal/storage"

type SingersStorage interface {
	GetSingerById(id uint64) (storage.Singer, bool)
}
