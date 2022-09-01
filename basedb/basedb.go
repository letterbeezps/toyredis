package basedb

import (
	"sync"

	"github.com/letterbeezps/toyredis/storage"
)

type (
	BaseDb struct {
		Mu sync.RWMutex

		HashStorage   storage.HashStorageInterface
		StringStorage storage.StringStorageInterface
	}
)

func NewBaseDb() *BaseDb {
	return &BaseDb{
		Mu:            sync.RWMutex{},
		HashStorage:   storage.NewHashStorage[[]byte](),
		StringStorage: storage.NewStringStorage[string](),
	}
}
