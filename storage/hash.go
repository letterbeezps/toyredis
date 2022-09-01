package storage

import (
	"fmt"
	"sync"

	"github.com/letterbeezps/toyredis/datastructure/hash"
)

type HashStorageInterface interface {
	HashHello()
}

type HashStorage[T any] struct {
	sync.RWMutex
	*hash.Hash[T]
}

func (hashStore *HashStorage[T]) HashHello() {
	fmt.Println("HashStorage")
}

func NewHashStorage[T any]() *HashStorage[T] {
	return &HashStorage[T]{
		Hash: hash.New[T](),
	}
}
