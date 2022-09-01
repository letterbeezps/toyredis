package storage

import (
	"fmt"
	"sync"

	"github.com/letterbeezps/toyredis/datastructure/trie"
)

type StringStorageInterface interface {
	StringHello()
	Get(key string) (isEnd bool, value string)
	Set(key, value string)
}

type StringStorage[T any] struct {
	sync.RWMutex
	*trie.Trie[T]
}

func (stringStore *StringStorage[T]) StringHello() {
	fmt.Println("StringStorage")
}

func (stringStore *StringStorage[T]) Get(key string) (isEnd bool, value T) {

	return stringStore.Search(key)
}

func (stringStore *StringStorage[T]) Set(key string, value T) {
	stringStore.Insert(key, value)
}

func NewStringStorage[T any]() *StringStorage[T] {
	return &StringStorage[T]{
		Trie: trie.New[T](),
	}
}
